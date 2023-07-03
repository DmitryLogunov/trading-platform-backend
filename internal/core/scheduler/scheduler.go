package scheduler

import (
	"errors"
	"fmt"
	cronPeriodUnits "github.com/DmitryLogunov/trading-platform-backend/internal/core/scheduler/enums/cron-period-units"
	jobsStatuses "github.com/DmitryLogunov/trading-platform-backend/internal/core/scheduler/enums/jobs-statuses"
	"log"
	"time"

	"github.com/go-co-op/gocron"
)

type JobsManager struct {
	jobs           *JobsStorage
	refreshJobsTag string
}

func (jm *JobsManager) Init() {
	jm.jobs = &JobsStorage{}
	jm.jobs.Init()

	s := gocron.NewScheduler(time.UTC)

	jm.refreshJobsTag = jm.jobs.AddJob(
		"scheduler-manager-refreshing-job",
		jm.RefreshJobs,
		"",
		CronPeriod{Unit: cronPeriodUnits.Seconds, Interval: 5},
	)

	_, err := s.Every(5).Seconds().Tag(jm.refreshJobsTag).Do(jm.RefreshJobs, "")
	if err != nil {
		log.Fatalln("error init scheduling manager", err)
	}

	jm.jobs.StartJob(jm.refreshJobsTag, s)
	s.StartAsync()

	log.Printf("Schedule jobs manager successfully started")
	//defer jm.terminateAllJobs()
}

func (jm *JobsManager) AddJob(
	handlerTag string,
	handler func(interface{}) bool,
	params interface{},
	cronPeriod CronPeriod,
) string {
	return jm.jobs.AddJob(handlerTag, handler, params, cronPeriod)
}

func (jm *JobsManager) DeleteJob(tag string) (bool, error) {
	if job := jm.jobs.ChangeJobStatus(tag, jobsStatuses.Finished); job == nil {
		return false, errors.New("JobsManager: finishing task failed")
	}

	return true, nil
}

func (jm *JobsManager) FindJobByTag(tag string) *Job {
	return jm.jobs.FindJobByTag(tag)
}

func (jm *JobsManager) FindAll() *map[string]*Job {
	return jm.jobs.FindAll()
}

func (jm *JobsManager) startJob(job *Job) (bool, error) {
	task := func(params interface{}) {
		job.handler(params)
		jm.jobs.RefreshProcessedAt(job.Tag)
	}

	s := gocron.NewScheduler(time.UTC)

	s.SingletonModeAll()

	cronConfiguredScheduler := jm.parseCronPeriod(s, job.CronPeriod)
	if cronConfiguredScheduler == nil {
		fmt.Printf("scheduler jobs manager error: wrong cron expression of job")
		return false, nil
	}

	_, err := cronConfiguredScheduler.Tag(job.Tag).Do(task, job.Params)
	if err != nil {
		log.Fatalln("error scheduling job", err)
		return false, err
	}

	jm.jobs.StartJob(job.Tag, s)
	s.StartAsync()

	log.Printf("Schedule job started: %s", job.HandlerTag)

	return true, nil
}

func (jm *JobsManager) RefreshJobs(interface{}) bool {
	for tag, job := range jm.jobs.data {
		if job.Status == jobsStatuses.Finished {
			job.scheduler.RemoveByTag(tag)
			jm.jobs.Delete(tag)
		}

		if job.Status == jobsStatuses.Created {
			jm.jobs.ChangeJobStatus(tag, jobsStatuses.Pending)
			jm.startJob(job)
		}
	}

	return true
}

func (jm *JobsManager) terminateAllJobs() {
	refreshJobsTask := jm.jobs.data[jm.refreshJobsTag]
	if refreshJobsTask != nil {
		refreshJobsTask.scheduler.RemoveByTag(jm.refreshJobsTag)
		jm.jobs.Delete(jm.refreshJobsTag)
	}

	for tag, job := range jm.jobs.data {
		if job.Status == jobsStatuses.InProcess {
			job.scheduler.RemoveByTag(tag)
		}

		jm.jobs.Delete(tag)
	}
}

func (jm *JobsManager) parseCronPeriod(s *gocron.Scheduler, cronPeriod CronPeriod) *gocron.Scheduler {
	if cronPeriod.Unit == cronPeriodUnits.Seconds {
		return s.Every(cronPeriod.Interval).Seconds()
	}

	if cronPeriod.Unit == cronPeriodUnits.Minutes {
		return s.Every(cronPeriod.Interval).Minutes()
	}

	if cronPeriod.Unit == cronPeriodUnits.Hours {
		return s.Every(cronPeriod.Interval).Hours()
	}

	return nil
}

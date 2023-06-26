package scheduler

import (
	"errors"
	"fmt"
	jobStatuses "github.com/DmitryLogunov/trading-platform/internal/core/scheduler/enums/jobs-statuses"
	"github.com/go-co-op/gocron"
	"github.com/google/uuid"
	"log"
	"time"
)

type JobsStorage struct {
	data map[string]*Job
}

// CronPeriod : repeat cron interval in unit (seconds, minutes, hours)
type CronPeriod struct {
	Unit     uint
	Interval int
}

type Job struct {
	Tag             string `bson:"tag,omitempty"`
	HandlerTag      string `bson:"handlerTag"`
	handler         func(interface{}) bool
	Params          interface{}       `bson:"handlerTag"`
	scheduler       *gocron.Scheduler `bson:"omitempty"`
	CronPeriod      CronPeriod
	CreatedAt       time.Time `bson:"createdAt"`
	UpdatedAt       time.Time `bson:"updatedAt,omitempty"`
	LastProcessedAt time.Time `bson:"lastProcessedAt,omitempty"`
	Status          uint      `bson:"status"`
}

func (ts *JobsStorage) Init() {
	ts.data = make(map[string]*Job)
}

func (ts *JobsStorage) AddJob(handlerTag string, handler func(interface{}) bool, params interface{}, cronPeriod CronPeriod) string {
	tag := uuid.New().String()

	ts.data[tag] = &Job{
		Tag:        tag,
		HandlerTag: handlerTag,
		handler:    handler,
		Params:     params,
		CronPeriod: CronPeriod{
			Unit:     cronPeriod.Unit,
			Interval: cronPeriod.Interval,
		},
		CreatedAt: time.Now(),
		Status:    jobStatuses.Created,
	}

	return tag
}

func (ts *JobsStorage) StartJob(tag string, scheduler *gocron.Scheduler) *Job {
	if ts.data[tag] == nil {
		return nil
	}

	t := ts.data[tag]
	ts.data[tag] = &Job{
		Tag:             t.Tag,
		HandlerTag:      t.HandlerTag,
		handler:         t.handler,
		Params:          t.Params,
		CronPeriod:      t.CronPeriod,
		scheduler:       scheduler,
		CreatedAt:       t.CreatedAt,
		UpdatedAt:       time.Now(),
		LastProcessedAt: t.LastProcessedAt,
		Status:          jobStatuses.InProcess,
	}

	return ts.data[tag]
}

func (ts *JobsStorage) FindJobByTag(tag string) *Job {
	return ts.data[tag]
}

func (ts *JobsStorage) FindAll() *map[string]*Job {
	return &ts.data
}

func (ts *JobsStorage) ChangeJobStatus(tag string, status uint) *Job {
	if ts.data[tag] == nil {
		return nil
	}

	t := ts.data[tag]
	ts.data[tag] = &Job{
		Tag:             t.Tag,
		HandlerTag:      t.HandlerTag,
		handler:         t.handler,
		Params:          t.Params,
		CronPeriod:      t.CronPeriod,
		scheduler:       t.scheduler,
		CreatedAt:       t.CreatedAt,
		UpdatedAt:       time.Now(),
		LastProcessedAt: t.LastProcessedAt,
		Status:          status,
	}

	return ts.data[tag]
}

func (ts *JobsStorage) RefreshProcessedAt(tag string) *Job {
	if ts.data[tag] == nil {
		return nil
	}

	t := ts.data[tag]
	ts.data[tag] = &Job{
		Tag:             t.Tag,
		HandlerTag:      t.HandlerTag,
		handler:         t.handler,
		Params:          t.Params,
		CronPeriod:      t.CronPeriod,
		scheduler:       t.scheduler,
		CreatedAt:       t.CreatedAt,
		UpdatedAt:       t.UpdatedAt,
		LastProcessedAt: time.Now(),
		Status:          t.Status,
	}

	return ts.data[tag]
}

func (ts *JobsStorage) Delete(tag string) (bool, error) {
	if ts.data[tag] == nil {
		fmt.Println("Error: Scheduler Jobs storage deleting: Job has not found")
		return false, errors.New("scheduler Jobs storage deleting: the Job not found")
	}

	delete(ts.data, tag)

	log.Printf("Schedule job terminated: %s", tag)

	return true, nil
}

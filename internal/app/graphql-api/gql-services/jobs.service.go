package gqlServices

import (
	"encoding/json"
	"fmt"
	graphqlApi "github.com/DmitryLogunov/trading-platform-backend/internal/app/graphql-api"
	"github.com/DmitryLogunov/trading-platform-backend/internal/core/scheduler"
	cronPeriodUnits "github.com/DmitryLogunov/trading-platform-backend/internal/core/scheduler/enums/cron-period-units"
	jobStatuses "github.com/DmitryLogunov/trading-platform-backend/internal/core/scheduler/enums/jobs-statuses"
	"github.com/DmitryLogunov/trading-platform-backend/internal/core/scheduler/handlers"
)

type JobsService struct{}

func (js *JobsService) StartJob(s *scheduler.JobsManager, input graphqlApi.JobData) (string, error) {
	var cronPeriodUnit uint
	if input.CronPeriod.Unit == "seconds" {
		cronPeriodUnit = cronPeriodUnits.Seconds
	}

	if input.CronPeriod.Unit == "minutes" {
		cronPeriodUnit = cronPeriodUnits.Minutes
	}

	if input.CronPeriod.Unit == "hours" {
		cronPeriodUnit = cronPeriodUnits.Hours
	}

	var handlerParams []handlers.HandlerParam
	for _, p := range input.Params {
		handlerParams = append(handlerParams, handlers.HandlerParam{
			Key:   p.Key,
			Value: p.Value,
		})
	}
	return s.AddJob(
		input.HandlerTag,
		handlerParams,
		scheduler.CronPeriod{Unit: cronPeriodUnit, Interval: input.CronPeriod.Interval},
	), nil
}

func (js *JobsService) StopJob(s *scheduler.JobsManager, tag string) (string, error) {
	if res, err := s.DeleteJob(tag); res == false || err != nil {
		return fmt.Sprintf("%s", err), err
	}

	return "OK", nil
}

func (js *JobsService) GetAllJobs(s *scheduler.JobsManager) ([]*graphqlApi.Job, error) {
	var gqlJobs []*graphqlApi.Job

	jobs := s.FindAll()

	if jobs == nil {
		return gqlJobs, nil
	}

	for _, job := range *jobs {
		paramsStringify, err := json.Marshal((*job).Params)
		if err != nil {
			panic(err)
		}

		cronPeriodStringify, err := json.Marshal((*job).CronPeriod)
		if err != nil {
			panic(err)
		}

		var statusStringify string
		if (*job).Status == jobStatuses.Created {
			statusStringify = "created"
		}

		if (*job).Status == jobStatuses.InProcess {
			statusStringify = "inProcess"
		}

		if (*job).Status == jobStatuses.Finished {
			statusStringify = "finished"
		}

		gqlJobs = append(gqlJobs, &graphqlApi.Job{
			Tag:        (*job).Tag,
			HandlerTag: (*job).HandlerTag,
			Params:     string(paramsStringify),
			CronPeriod: string(cronPeriodStringify),
			CreatedAt:  (*job).CreatedAt,
			Status:     statusStringify,
		})
	}

	return gqlJobs, nil
}

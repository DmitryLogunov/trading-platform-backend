package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.33

import (
	"context"
	"encoding/json"
	"fmt"
	jobStatuses "github.com/DmitryLogunov/trading-platform/internal/core/scheduler/enums/jobs-statuses"

	graphql_api "github.com/DmitryLogunov/trading-platform/internal/api/graphql-api"
	mongodbModels "github.com/DmitryLogunov/trading-platform/internal/database/mongodb/models"
)

// GetPosts is the resolver for the getPosts field.
func (r *queryResolver) GetPosts(ctx context.Context) ([]*graphql_api.Post, error) {
	postsModel := mongodbModels.Post{}

	posts, err := postsModel.GetPosts(ctx, r.MongoDB)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var gqlPosts []*graphql_api.Post

	for _, p := range posts {
		gqlPosts = append(gqlPosts, &graphql_api.Post{
			ID:          p.ID.Hex(),
			Title:       p.Title,
			Content:     p.Content,
			Author:      p.Author,
			Hero:        p.Hero,
			PublishedAt: p.PublishedAt,
			UpdatedAt:   p.UpdatedAt,
		})
	}

	return gqlPosts, nil
}

// GetAllJobs is the resolver for the getAllJobs field.
func (r *queryResolver) GetAllJobs(ctx context.Context) ([]*graphql_api.Job, error) {
	var gqlJobs []*graphql_api.Job

	jobs := r.Scheduler.FindAll()

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

		gqlJobs = append(gqlJobs, &graphql_api.Job{
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

// Query returns graphql_api.QueryResolver implementation.
func (r *Resolver) Query() graphql_api.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

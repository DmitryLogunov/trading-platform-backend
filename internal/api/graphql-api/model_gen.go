// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphql_api

import (
	"time"
)

type CronPeriodInput struct {
	Unit     string `json:"unit"`
	Interval int    `json:"interval"`
}

type CronPeriodObject struct {
	Unit     string `json:"unit"`
	Interval int    `json:"interval"`
}

type Job struct {
	Tag             string     `json:"tag"`
	HandlerTag      string     `json:"handlerTag"`
	Params          string     `json:"params"`
	CronPeriod      string     `json:"cronPeriod"`
	CreatedAt       time.Time  `json:"createdAt"`
	UpdatedAt       *time.Time `json:"updatedAt,omitempty"`
	LastProcessedAt *time.Time `json:"lastProcessedAt,omitempty"`
	Status          string     `json:"status"`
}

type JobData struct {
	HandlerTag string           `json:"handlerTag"`
	Params     []*JobParamInput `json:"params"`
	CronPeriod *CronPeriodInput `json:"cronPeriod"`
}

type JobParamInput struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}

type JobParamObject struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}

type NewPost struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
	Hero    string `json:"hero"`
}

type Post struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Author      string    `json:"author"`
	Hero        string    `json:"hero"`
	PublishedAt time.Time `json:"publishedAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

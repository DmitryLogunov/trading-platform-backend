package gqlServices

import (
	"context"
	"fmt"
	graphqlApi "github.com/DmitryLogunov/trading-platform/internal/app/graphql-api"
	mongodbModels "github.com/DmitryLogunov/trading-platform/internal/core/database/mongodb/models"
	"github.com/DmitryLogunov/trading-platform/internal/core/helpers"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type AlertsService struct{}

// GetAlerts : returns the list of all tradings from DB
func (as *AlertsService) GetAlerts(ctx context.Context, mongoDB *mongo.Database, filters *graphqlApi.AlertsFiltersInput) ([]*graphqlApi.Alert, error) {
	mongodbFilters, err := as.parseFilters(filters)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	alertsModel := mongodbModels.Alert{}

	alerts, err := alertsModel.Find(ctx, mongoDB, mongodbFilters)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var gqlAlerts []*graphqlApi.Alert

	for _, a := range alerts {
		gqlAlerts = append(gqlAlerts, &graphqlApi.Alert{
			ID:        a.ID.Hex(),
			Title:     a.Title,
			Ticker:    a.Ticker,
			Action:    int(a.Action),
			Price:     float64(a.Price),
			CreatedAt: a.CreatedAt,
		})
	}

	return gqlAlerts, nil
}

// parseFilters : parses Graphql input filters to MongoDB filters
func (as *AlertsService) parseFilters(filters *graphqlApi.AlertsFiltersInput) (*mongodbModels.AlertsFilters, error) {
	var err error
	var createdAtFrom *time.Time
	var createdAtTo *time.Time

	if filters == nil || filters.CreatedAtFrom == nil || *filters.CreatedAtFrom == "" {
		unixStartTime := time.Unix(0, 0).UTC()
		createdAtFrom = &unixStartTime
	} else if filters != nil && filters.CreatedAtFrom != nil {
		createdAtFrom, err = helpers.DatetimeParse(*filters.CreatedAtFrom)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}

	if filters == nil || filters.CreatedAtTo == nil || *filters.CreatedAtTo == "" {
		timeNow := time.Now().UTC()
		createdAtTo = &timeNow
	} else if filters != nil && filters.CreatedAtTo != nil {
		createdAtTo, err = helpers.DatetimeParse(*filters.CreatedAtTo)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}

	title := ""
	if filters != nil && filters.Title != nil {
		title = *filters.Title
	}

	ticker := ""
	if filters != nil && filters.Ticker != nil {
		ticker = *filters.Ticker
	}

	action := uint(2)
	if filters != nil && filters.Action != nil && *filters.Action == "buy" {
		action = 0
	} else if filters != nil && filters.Action != nil && *filters.Action == "sell" {
		action = 1
	}

	return &mongodbModels.AlertsFilters{
		Title:  title,
		Ticker: ticker,
		Action: action,
		CreatedAt: &mongodbModels.DatetimeComparingFilter{
			From: createdAtFrom,
			To:   createdAtTo,
		},
	}, nil
}

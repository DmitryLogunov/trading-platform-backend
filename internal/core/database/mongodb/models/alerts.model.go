package models

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Alert struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Title     string             `bson:"title",gorm:"not null"`
	Ticker    string             `bson:"ticker",gorm:"not null"`
	Action    uint               `bson:"action",gorm:"not null"`
	Price     float32            `bson:"price",gorm:"not null"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
}

type DatetimeComparingFilter struct {
	start time.Time
	end   time.Time
}

type AlertsFilters struct {
	Title     string
	Ticker    string
	Action    uint
	CreatedAt DatetimeComparingFilter
}

// getCollection: returns "posts" mongodb collection
func (a *Alert) getCollection(db *mongo.Database) *mongo.Collection {
	collectionName := "alerts"

	return db.Collection(collectionName)
}

// Save : saves alert in db
func (a *Alert) Save(ctx context.Context, db *mongo.Database, input *Alert) (*Alert, error) {
	alert := Alert{
		Title:     input.Title,
		Ticker:    input.Ticker,
		Action:    input.Action,
		Price:     input.Price,
		CreatedAt: input.CreatedAt,
	}

	res, err := a.getCollection(db).InsertOne(ctx, alert)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	alert.ID = res.InsertedID.(primitive.ObjectID)

	return &alert, nil
}

// Find : returns alerts using filters
func (a *Alert) Find(ctx context.Context, db *mongo.Database, filters *AlertsFilters) ([]*Alert, error) {
	cursor, err := a.getCollection(db).Find(ctx, bson.M{})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var alerts []*Alert
	for cursor.Next(context.Background()) {
		alert := Alert{}
		cursor.Decode(&alert)

		alerts = append(alerts, &alert)
	}

	return alerts, nil
}

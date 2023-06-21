package models

import (
	"context"
	"fmt"
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

// getCollection: returns "posts" mongodb collection
func (a *Alert) getCollection(db *mongo.Database) *mongo.Collection {
	collectionName := "alerts"

	return db.Collection(collectionName)
}

// SaveAlert : saves alert in db
func (a *Alert) SaveAlert(ctx context.Context, db *mongo.Database, input *Alert) (*Alert, error) {
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

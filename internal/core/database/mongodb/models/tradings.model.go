package models

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Trading struct {
	ID                                primitive.ObjectID `bson:"_id,omitempty"`
	Exchange                          string             `bson:"exchange",gorm:"not null"`
	BaseCurrency                      string             `bson:"base_currency",gorm:"not null"`
	SecondaryCurrency                 string             `bson:"secondary_currency",gorm:"not null"`
	BaseDepositInBaseCurrency         float32            `bson:"base_deposit_in_base_currency",gorm:"not null; unique"`
	CurrentDepositInBaseCurrency      float32            `bson:"current_deposit_in_base_currency,omitempty"`
	CurrentDepositInSecondaryCurrency float32            `bson:"current_deposit_in_secondary_currency,omitempty"`
	RoiInPercent                      float32            `bson:"roi_in_percent,omitempty"`
	RoiInBaseCurrency                 float32            `bson:"roi_in_base_currency,omitempty"`
	StartedAt                         time.Time          `bson:"started_at"`
	ClosedAt                          time.Time          `bson:"closed_at,omitempty"`
}

// getCollection: returns "posts" mongodb collection
func (t *Trading) getCollection(db *mongo.Database) *mongo.Collection {
	collectionName := "tradings"

	return db.Collection(collectionName)
}

// CreateTrading : creates posts
func (t *Trading) CreateTrading(ctx context.Context, db *mongo.Database, input *Trading) (*Trading, error) {
	newTrading := Trading{
		Exchange:                     input.Exchange,
		BaseCurrency:                 input.BaseCurrency,
		SecondaryCurrency:            input.SecondaryCurrency,
		BaseDepositInBaseCurrency:    input.BaseDepositInBaseCurrency,
		CurrentDepositInBaseCurrency: input.BaseDepositInBaseCurrency,
		StartedAt:                    input.StartedAt,
	}

	res, err := t.getCollection(db).InsertOne(ctx, newTrading)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	newTrading.ID = res.InsertedID.(primitive.ObjectID)

	return &newTrading, nil
}

// GetTradings : returns posts
func (t *Trading) GetTradings(ctx context.Context, db *mongo.Database) ([]*Trading, error) {
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"started_at", -1}})

	cursor, err := t.getCollection(db).Find(ctx, bson.M{}, findOptions)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var tradings []*Trading
	for cursor.Next(context.Background()) {
		trading := Trading{}
		cursor.Decode(&trading)

		tradings = append(tradings, &trading)
	}

	return tradings, nil
}

// DeleteTrading : deletes trading by ID
func (t *Trading) DeleteTrading(ctx context.Context, db *mongo.Database, id string) (bool, error) {
	parsedId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	res, err := t.getCollection(db).DeleteOne(ctx, bson.D{{"_id", parsedId}})
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	if res.DeletedCount == 0 {
		return false, fmt.Errorf("not found")
	}

	return true, nil
}

// GetTradingByID : returns trading by ID
func (t *Trading) GetTradingByID(ctx context.Context, db *mongo.Database, id string) (*Trading, error) {
	parsedId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	trading := Trading{}
	err = t.getCollection(db).FindOne(ctx, bson.D{{"_id", parsedId}}).Decode(&trading)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &trading, nil
}

// UpdateTrading : updates trading by ID
func (t *Trading) UpdateTrading(ctx context.Context, db *mongo.Database, id string, input *Trading) (*Trading, error) {
	parsedId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	updateTradingBsonData := bson.D{{
		"$set",
		bson.D{
			{"base_deposit_in_base_currency", input.BaseDepositInBaseCurrency},
			{"current_deposit_in_base_currency", input.CurrentDepositInBaseCurrency},
			{"current_deposit_in_secondary_currency", input.CurrentDepositInSecondaryCurrency},
			{"roi_in_percent", input.RoiInPercent},
			{"roi_in_base_currency", input.RoiInBaseCurrency},
			{"closed_at", input.ClosedAt},
		}}}

	resMainRequest, err := t.getCollection(db).UpdateOne(ctx, bson.D{{"_id", parsedId}}, updateTradingBsonData)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	resClosedRefreshModifiedCount := int64(0)
	if input.ClosedAt == time.Unix(0, 0).UTC() {
		updateTradingBsonData := bson.D{{
			"$unset",
			bson.D{
				{"closed_at", ""},
			}}}

		resClosedRefreshRequest, err := t.getCollection(db).UpdateOne(ctx, bson.D{{"_id", parsedId}}, updateTradingBsonData)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		resClosedRefreshModifiedCount = resClosedRefreshRequest.ModifiedCount
	}

	if resMainRequest.ModifiedCount+resClosedRefreshModifiedCount < 1 {
		return nil, fmt.Errorf("updating failed")
	}

	return t.GetTradingByID(ctx, db, id)
}

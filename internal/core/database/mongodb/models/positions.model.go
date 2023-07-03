package models

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"

	actionsEnum "github.com/DmitryLogunov/trading-platform-backend/internal/core/database/mongodb/enums/actions"
)

type Order struct {
	Action               uint       `bson:"action",gorm:"not null"`
	SourceCurrencyAmount float32    `bson:"source_currency_amount",gorm:"not null"`
	TargetCurrencyAmount float32    `bson:"target_currency_amount",gorm:"not null"`
	Price                float32    `bson:"price",gorm:"not null"`
	CreatedAt            *time.Time `bson:"created_at"`
}

type OpenPositionData struct {
	TradingId          string     `bson:"trading_id",gorm:"not null"`
	BaseCurrencyAmount float32    `bson:"source_currency_amount",gorm:"not null"`
	Price              float32    `bson:"price",gorm:"not null"`
	CreatedAt          *time.Time `bson:"created_at"`
}

type ClosePositionData struct {
	ID       string     `bson:"id",gorm:"not null"`
	Price    float32    `bson:"price",gorm:"not null"`
	ClosedAt *time.Time `bson:"closed_at"`
}

type Position struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	TradingId         string             `bson:"trading_id",gorm:"not null"`
	BaseCurrency      string             `bson:"base_currency",gorm:"not null"`
	SecondaryCurrency string             `bson:"second_currency",gorm:"not null"`
	Orders            map[uint]*Order    `bson:"orders",gorm:"not null"`
	RoiInPercent      float32            `bson:"roi_in_percent,omitempty"`
	RoiInBaseCurrency float32            `bson:"roi_in_base_currency,omitempty"`
	CreatedAt         *time.Time         `bson:"created_at"`
	ClosedAt          *time.Time         `bson:"closed_at,omitempty"`
}

// getCollection: returns "posts" mongodb collection
func (p *Position) getCollection(db *mongo.Database) *mongo.Collection {
	collectionName := "positions"

	return db.Collection(collectionName)
}

// OpenPosition : created and saves position item on MongoDB
func (p *Position) OpenPosition(ctx context.Context, db *mongo.Database, input *OpenPositionData) (*Position, error) {
	tradingModelItem := &Trading{}
	trading, err := tradingModelItem.GetTradingByID(ctx, db, input.TradingId)
	if err != nil || trading == nil {
		return nil, fmt.Errorf("trading not found")
	}

	if input.BaseCurrencyAmount == 0 {
		return nil, fmt.Errorf("base currency amount should not be 0")
	}

	if input.Price == 0 {
		return nil, fmt.Errorf("price should not be 0")
	}

	openPositionOrder := &Order{
		Action:               actionsEnum.Buy,
		SourceCurrencyAmount: input.BaseCurrencyAmount,
		TargetCurrencyAmount: input.BaseCurrencyAmount / input.Price,
		Price:                input.Price,
		CreatedAt:            input.CreatedAt,
	}

	position := Position{
		TradingId:         input.TradingId,
		BaseCurrency:      trading.BaseCurrency,
		SecondaryCurrency: trading.SecondaryCurrency,
		Orders:            map[uint]*Order{actionsEnum.Buy: openPositionOrder},
		CreatedAt:         input.CreatedAt,
	}

	res, err := p.getCollection(db).InsertOne(ctx, position)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	position.ID = res.InsertedID.(primitive.ObjectID)

	return &position, nil
}

// ClosePosition : created and saves position item on MongoDB
func (p *Position) ClosePosition(ctx context.Context, db *mongo.Database, input *ClosePositionData) (*Position, error) {
	parsedId, err := primitive.ObjectIDFromHex(input.ID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if input.Price == 0 {
		return nil, fmt.Errorf("price should not be 0")
	}

	position, err := p.GetPositionByID(ctx, db, input.ID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	startBaseCurrencyAmount := position.Orders[actionsEnum.Buy].SourceCurrencyAmount
	finalBaseCurrencyAmount := position.Orders[actionsEnum.Buy].TargetCurrencyAmount * input.Price

	closePositionOrder := &Order{
		Action:               actionsEnum.Sell,
		SourceCurrencyAmount: position.Orders[actionsEnum.Buy].TargetCurrencyAmount,
		TargetCurrencyAmount: finalBaseCurrencyAmount,
		Price:                input.Price,
		CreatedAt:            input.ClosedAt,
	}

	refreshedOrders := map[uint]*Order{
		actionsEnum.Buy:  position.Orders[actionsEnum.Buy],
		actionsEnum.Sell: closePositionOrder,
	}

	roiInPercent := (finalBaseCurrencyAmount - startBaseCurrencyAmount) / startBaseCurrencyAmount
	roiInBaseCurrency := finalBaseCurrencyAmount - startBaseCurrencyAmount

	updatePositionBsonData := bson.D{{
		"$set",
		bson.D{
			{"orders", refreshedOrders},
			{"roi_in_percent", roiInPercent},
			{"roi_in_base_currency", roiInBaseCurrency},
			{"closed_at", input.ClosedAt},
		}}}

	_, err = p.getCollection(db).UpdateOne(ctx, bson.D{{"_id", parsedId}}, updatePositionBsonData)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return p.GetPositionByID(ctx, db, input.ID)
}

// DeletePosition : deletes position by ID
func (p *Position) DeletePosition(ctx context.Context, db *mongo.Database, id string) (bool, error) {
	parsedId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	res, err := p.getCollection(db).DeleteOne(ctx, bson.D{{"_id", parsedId}})
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	if res.DeletedCount == 0 {
		return false, fmt.Errorf("not found")
	}

	return true, nil
}

// GetPositionByID : returns position by ID
func (p *Position) GetPositionByID(ctx context.Context, db *mongo.Database, id string) (*Position, error) {
	parsedId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	position := Position{}
	err = p.getCollection(db).FindOne(ctx, bson.D{{"_id", parsedId}}).Decode(&position)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &position, nil
}

// Find : returns alerts using filters
func (p *Position) Find(ctx context.Context, db *mongo.Database, TradingId string) ([]*Position, error) {
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"created_at", -1}})

	cursor, err := p.getCollection(db).Find(ctx, bson.D{{"trading_id", TradingId}}, findOptions)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var positions []*Position
	for cursor.Next(context.Background()) {
		position := Position{}
		cursor.Decode(&position)

		positions = append(positions, &position)
	}

	return positions, nil
}

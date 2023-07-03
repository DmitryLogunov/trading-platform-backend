package gqlServices

import (
	"context"
	"fmt"
	graphqlApi "github.com/DmitryLogunov/trading-platform-backend/internal/app/graphql-api"
	"github.com/DmitryLogunov/trading-platform-backend/internal/core/database/mongodb/enums/actions"
	mongodbModels "github.com/DmitryLogunov/trading-platform-backend/internal/core/database/mongodb/models"
	"github.com/DmitryLogunov/trading-platform-backend/internal/core/helpers"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type PositionsService struct{}

// OpenPosition : creates and saves new position into DB
func (ps *PositionsService) OpenPosition(ctx context.Context, mongoDB *mongo.Database, input graphqlApi.OpenPositionInput) (*graphqlApi.Position, error) {
	timeNow := time.Now()
	var createdAt *time.Time = &timeNow
	var err error

	if input.CreatedAt != nil {
		createdAt, err = helpers.DatetimeParse(*input.CreatedAt)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}

	openPositionData := mongodbModels.OpenPositionData{
		TradingID:          input.TradingID,
		BaseCurrencyAmount: float32(input.BaseCurrencyAmount),
		Price:              float32(input.Price),
		CreatedAt:          createdAt,
	}

	positionModelItem := &mongodbModels.Position{}
	newPosition, err := positionModelItem.OpenPosition(ctx, mongoDB, &openPositionData)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	gqlBuyOrder := &graphqlApi.Order{
		Action:               int(newPosition.Orders[actions.Buy].Action),
		SourceCurrencyAmount: float64(newPosition.Orders[actions.Buy].SourceCurrencyAmount),
		TargetCurrencyAmount: float64(newPosition.Orders[actions.Buy].TargetCurrencyAmount),
		Price:                float64(newPosition.Orders[actions.Buy].Price),
		CreatedAt:            *newPosition.Orders[actions.Buy].CreatedAt,
	}

	return &graphqlApi.Position{
		ID:                newPosition.ID.Hex(),
		TradingID:         newPosition.TradingID,
		BaseCurrency:      newPosition.BaseCurrency,
		SecondaryCurrency: newPosition.SecondaryCurrency,
		Orders:            []*graphqlApi.Order{gqlBuyOrder},
		CreatedAt:         *newPosition.CreatedAt,
	}, nil
}

// ClosePosition : closes position in DB
func (ps *PositionsService) ClosePosition(ctx context.Context, mongoDB *mongo.Database, input graphqlApi.ClosePositionInput) (*graphqlApi.Position, error) {
	timeNow := time.Now()
	var closedAt *time.Time = &timeNow
	var err error

	if input.ClosedAt != nil {
		closedAt, err = helpers.DatetimeParse(*input.ClosedAt)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}

	closePositionData := mongodbModels.ClosePositionData{
		ID:       input.ID,
		Price:    float32(input.Price),
		ClosedAt: closedAt,
	}

	positionModelItem := &mongodbModels.Position{}
	closedPosition, err := positionModelItem.ClosePosition(ctx, mongoDB, &closePositionData)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	gqlBuyOrder := &graphqlApi.Order{
		Action:               int(closedPosition.Orders[actions.Buy].Action),
		SourceCurrencyAmount: float64(closedPosition.Orders[actions.Buy].SourceCurrencyAmount),
		TargetCurrencyAmount: float64(closedPosition.Orders[actions.Buy].TargetCurrencyAmount),
		Price:                float64(closedPosition.Orders[actions.Buy].Price),
		CreatedAt:            *closedPosition.Orders[actions.Buy].CreatedAt,
	}

	gqlSellOrder := &graphqlApi.Order{
		Action:               int(closedPosition.Orders[actions.Sell].Action),
		SourceCurrencyAmount: float64(closedPosition.Orders[actions.Sell].SourceCurrencyAmount),
		TargetCurrencyAmount: float64(closedPosition.Orders[actions.Sell].TargetCurrencyAmount),
		Price:                float64(closedPosition.Orders[actions.Sell].Price),
		CreatedAt:            *closedPosition.Orders[actions.Sell].CreatedAt,
	}

	roiInPercent := float64(closedPosition.RoiInPercent)
	roiInBaseCurrency := float64(closedPosition.RoiInBaseCurrency)

	return &graphqlApi.Position{
		ID:                closedPosition.ID.Hex(),
		TradingID:         closedPosition.TradingID,
		BaseCurrency:      closedPosition.BaseCurrency,
		SecondaryCurrency: closedPosition.SecondaryCurrency,
		Orders:            []*graphqlApi.Order{gqlBuyOrder, gqlSellOrder},
		RoiInPercent:      &roiInPercent,
		RoiInBaseCurrency: &roiInBaseCurrency,
		CreatedAt:         *closedPosition.CreatedAt,
		ClosedAt:          closedPosition.ClosedAt,
	}, nil
}

// DeletePosition : deletes position by ID
func (ps *PositionsService) DeletePosition(ctx context.Context, mongoDB *mongo.Database, id string) (bool, error) {
	positionModelItem := mongodbModels.Position{}

	return positionModelItem.DeletePosition(ctx, mongoDB, id)
}

// GetPositions : returns the list of all positions from DB by tradingID
func (ps *PositionsService) GetPositions(ctx context.Context, mongoDB *mongo.Database, tradingID string) ([]*graphqlApi.Position, error) {
	positionsModelItem := mongodbModels.Position{}

	positions, err := positionsModelItem.Find(ctx, mongoDB, tradingID)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var gqlPositions []*graphqlApi.Position
	for _, p := range positions {
		gqlBuyOrder := &graphqlApi.Order{
			Action:               int(p.Orders[actions.Buy].Action),
			SourceCurrencyAmount: float64(p.Orders[actions.Buy].SourceCurrencyAmount),
			TargetCurrencyAmount: float64(p.Orders[actions.Buy].TargetCurrencyAmount),
			Price:                float64(p.Orders[actions.Buy].Price),
			CreatedAt:            *p.Orders[actions.Buy].CreatedAt,
		}

		orders := []*graphqlApi.Order{gqlBuyOrder}

		if p.Orders[actions.Sell] != nil {
			gqlSellOrder := &graphqlApi.Order{
				Action:               int(p.Orders[actions.Sell].Action),
				SourceCurrencyAmount: float64(p.Orders[actions.Sell].SourceCurrencyAmount),
				TargetCurrencyAmount: float64(p.Orders[actions.Sell].TargetCurrencyAmount),
				Price:                float64(p.Orders[actions.Sell].Price),
				CreatedAt:            *p.Orders[actions.Sell].CreatedAt,
			}

			orders = append(orders, gqlSellOrder)
		}

		roiInPercent := float64(p.RoiInPercent)
		roiInBaseCurrency := float64(p.RoiInBaseCurrency)

		gqlPositions = append(gqlPositions, &graphqlApi.Position{
			ID:                p.ID.Hex(),
			TradingID:         p.TradingID,
			BaseCurrency:      p.BaseCurrency,
			SecondaryCurrency: p.SecondaryCurrency,
			Orders:            orders,
			RoiInPercent:      &roiInPercent,
			RoiInBaseCurrency: &roiInBaseCurrency,
			CreatedAt:         *p.CreatedAt,
			ClosedAt:          p.ClosedAt,
		})
	}

	return gqlPositions, nil
}

// GetPositionByID : returns position by ID
func (ps *PositionsService) GetPositionByID(ctx context.Context, mongoDB *mongo.Database, id string) (*graphqlApi.Position, error) {
	positionsModelItem := mongodbModels.Position{}

	p, err := positionsModelItem.GetPositionByID(ctx, mongoDB, id)
	if err != nil {
		return nil, err
	}

	gqlBuyOrder := &graphqlApi.Order{
		Action:               int(p.Orders[actions.Buy].Action),
		SourceCurrencyAmount: float64(p.Orders[actions.Buy].SourceCurrencyAmount),
		TargetCurrencyAmount: float64(p.Orders[actions.Buy].TargetCurrencyAmount),
		Price:                float64(p.Orders[actions.Buy].Price),
		CreatedAt:            *p.Orders[actions.Buy].CreatedAt,
	}

	orders := []*graphqlApi.Order{gqlBuyOrder}

	if p.Orders[actions.Sell] != nil {
		gqlSellOrder := &graphqlApi.Order{
			Action:               int(p.Orders[actions.Sell].Action),
			SourceCurrencyAmount: float64(p.Orders[actions.Sell].SourceCurrencyAmount),
			TargetCurrencyAmount: float64(p.Orders[actions.Sell].TargetCurrencyAmount),
			Price:                float64(p.Orders[actions.Sell].Price),
			CreatedAt:            *p.Orders[actions.Sell].CreatedAt,
		}

		orders = append(orders, gqlSellOrder)
	}

	roiInPercent := float64(p.RoiInPercent)
	roiInBaseCurrency := float64(p.RoiInBaseCurrency)

	return &graphqlApi.Position{
		ID:                p.ID.Hex(),
		TradingID:         p.TradingID,
		BaseCurrency:      p.BaseCurrency,
		SecondaryCurrency: p.SecondaryCurrency,
		Orders:            orders,
		RoiInPercent:      &roiInPercent,
		RoiInBaseCurrency: &roiInBaseCurrency,
		CreatedAt:         *p.CreatedAt,
		ClosedAt:          p.ClosedAt,
	}, nil
}

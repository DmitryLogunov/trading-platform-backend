package gqlServices

import (
	"context"
	"fmt"
	graphqlApi "github.com/DmitryLogunov/trading-platform-backend/internal/app/graphql-api"
	actionsEnum "github.com/DmitryLogunov/trading-platform-backend/internal/core/database/mongodb/enums/actions"
	mongodbModels "github.com/DmitryLogunov/trading-platform-backend/internal/core/database/mongodb/models"
	"github.com/DmitryLogunov/trading-platform-backend/internal/core/helpers"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type TradingService struct{}

// CreateTrading : creates and saves new trading into DB
func (ts *TradingService) CreateTrading(ctx context.Context, mongoDB *mongo.Database, input graphqlApi.NewTradingInput) (*graphqlApi.Trading, error) {
	timeNow := time.Now()
	var startedAt *time.Time = &timeNow
	var err error

	if input.StartedAt != nil {
		startedAt, err = helpers.DatetimeParse(*input.StartedAt)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}

	tradingsModelItem := mongodbModels.Trading{
		Exchange:                  input.Exchange,
		BaseCurrency:              input.BaseCurrency,
		SecondaryCurrency:         input.SecondaryCurrency,
		BaseDepositInBaseCurrency: float32(input.BaseDepositInBaseCurrency),
		StartedAt:                 *startedAt,
	}

	newTrading, err := tradingsModelItem.CreateTrading(ctx, mongoDB, &tradingsModelItem)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &graphqlApi.Trading{
		ID:                        newTrading.ID.Hex(),
		Exchange:                  newTrading.Exchange,
		BaseCurrency:              newTrading.BaseCurrency,
		SecondaryCurrency:         newTrading.SecondaryCurrency,
		BaseDepositInBaseCurrency: float64(newTrading.BaseDepositInBaseCurrency),
		StartedAt:                 newTrading.StartedAt,
	}, nil
}

// DeleteTrading : deletes tradings by ID
func (ts *TradingService) DeleteTrading(ctx context.Context, mongoDB *mongo.Database, id string) (bool, error) {
	tradingsModelItem := mongodbModels.Trading{}

	return tradingsModelItem.DeleteTrading(ctx, mongoDB, id)
}

// UpdateTrading : updates trading in DB
func (ts *TradingService) UpdateTrading(ctx context.Context, mongoDB *mongo.Database, input graphqlApi.UpdateTradingInput) (*graphqlApi.Trading, error) {
	gqlTradingFromDB, err := ts.GetTradingByID(ctx, mongoDB, input.ID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var closedAt *time.Time
	unixStartDatetime := time.Unix(0, 0).UTC()
	if input.ClosedAt != nil && *input.ClosedAt == "null" {
		closedAt = &unixStartDatetime
	} else if input.ClosedAt != nil {
		closedAt, err = helpers.DatetimeParse(*input.ClosedAt)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	} else if gqlTradingFromDB.ClosedAt != nil {
		closedAt = gqlTradingFromDB.ClosedAt
	} else {
		closedAt = &unixStartDatetime
	}

	tradingsModelItem := mongodbModels.Trading{
		BaseDepositInBaseCurrency:         *helpers.ValueOrNil32(gqlTradingFromDB.BaseDepositInBaseCurrency, input.BaseDepositInBaseCurrency),
		CurrentDepositInBaseCurrency:      *helpers.ValueOrNil32(*gqlTradingFromDB.CurrentDepositInBaseCurrency, input.CurrentDepositInBaseCurrency),
		CurrentDepositInSecondaryCurrency: *helpers.ValueOrNil32(*gqlTradingFromDB.CurrentDepositInSecondaryCurrency, input.CurrentDepositInSecondaryCurrency),
		RoiInPercent:                      *helpers.ValueOrNil32(*gqlTradingFromDB.RoiInPercent, input.RoiInPercent),
		RoiInBaseCurrency:                 *helpers.ValueOrNil32(*gqlTradingFromDB.RoiInBaseCurrency, input.RoiInBaseCurrency),
		ClosedAt:                          *closedAt,
	}

	updatedTrading, err := tradingsModelItem.UpdateTrading(ctx, mongoDB, input.ID, &tradingsModelItem)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &graphqlApi.Trading{
		ID:                                updatedTrading.ID.Hex(),
		Exchange:                          updatedTrading.Exchange,
		BaseCurrency:                      updatedTrading.BaseCurrency,
		SecondaryCurrency:                 updatedTrading.SecondaryCurrency,
		BaseDepositInBaseCurrency:         float64(updatedTrading.BaseDepositInBaseCurrency),
		CurrentDepositInBaseCurrency:      helpers.SafeParseFloat64(&updatedTrading.CurrentDepositInBaseCurrency),
		CurrentDepositInSecondaryCurrency: helpers.SafeParseFloat64(&updatedTrading.CurrentDepositInSecondaryCurrency),
		RoiInPercent:                      helpers.SafeParseFloat64(&updatedTrading.RoiInPercent),
		RoiInBaseCurrency:                 helpers.SafeParseFloat64(&updatedTrading.RoiInBaseCurrency),
		StartedAt:                         updatedTrading.StartedAt,
		ClosedAt:                          &updatedTrading.ClosedAt,
	}, nil
}

// GetTradings : returns the list of all tradings from DB
func (ts *TradingService) GetTradings(ctx context.Context, mongoDB *mongo.Database) ([]*graphqlApi.Trading, error) {
	tradingsModelItem := mongodbModels.Trading{}

	tradings, err := tradingsModelItem.GetTradings(ctx, mongoDB)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var gqlTradings []*graphqlApi.Trading
	for _, t := range tradings {
		gqlTradings = append(gqlTradings, &graphqlApi.Trading{
			ID:                                t.ID.Hex(),
			Exchange:                          t.Exchange,
			BaseCurrency:                      t.BaseCurrency,
			SecondaryCurrency:                 t.SecondaryCurrency,
			BaseDepositInBaseCurrency:         float64(t.BaseDepositInBaseCurrency),
			CurrentDepositInBaseCurrency:      helpers.SafeParseFloat64(&t.CurrentDepositInBaseCurrency),
			CurrentDepositInSecondaryCurrency: helpers.SafeParseFloat64(&t.CurrentDepositInSecondaryCurrency),
			RoiInPercent:                      helpers.SafeParseFloat64(&t.RoiInPercent),
			RoiInBaseCurrency:                 helpers.SafeParseFloat64(&t.RoiInBaseCurrency),
			StartedAt:                         t.StartedAt,
			ClosedAt:                          &t.ClosedAt,
		})
	}

	return gqlTradings, nil
}

// GetTradingByID : returns trading by ID
func (ts *TradingService) GetTradingByID(ctx context.Context, mongoDB *mongo.Database, id string) (*graphqlApi.Trading, error) {
	tradingsModelItem := mongodbModels.Trading{}

	foundTrading, err := tradingsModelItem.GetTradingByID(ctx, mongoDB, id)
	if err != nil {
		return nil, err
	}

	gqlTrading := &graphqlApi.Trading{
		ID:                                foundTrading.ID.Hex(),
		Exchange:                          foundTrading.Exchange,
		BaseCurrency:                      foundTrading.BaseCurrency,
		SecondaryCurrency:                 foundTrading.SecondaryCurrency,
		BaseDepositInBaseCurrency:         float64(foundTrading.BaseDepositInBaseCurrency),
		CurrentDepositInBaseCurrency:      helpers.SafeParseFloat64(&foundTrading.CurrentDepositInBaseCurrency),
		CurrentDepositInSecondaryCurrency: helpers.SafeParseFloat64(&foundTrading.CurrentDepositInSecondaryCurrency),
		RoiInPercent:                      helpers.SafeParseFloat64(&foundTrading.RoiInPercent),
		RoiInBaseCurrency:                 helpers.SafeParseFloat64(&foundTrading.RoiInBaseCurrency),
		StartedAt:                         foundTrading.StartedAt,
		ClosedAt:                          &foundTrading.ClosedAt,
	}

	return gqlTrading, nil
}

// RefreshTrading : refreshes trading ROI and current balances
func (ts *TradingService) RefreshTrading(ctx context.Context, mongoDB *mongo.Database, id string) (*graphqlApi.Trading, error) {
	trading, err := ts.GetTradingByID(ctx, mongoDB, id)
	if err != nil {
		return nil, err
	}

	positionsModelItem := &mongodbModels.Position{}
	tradingPositions, err := positionsModelItem.Find(ctx, mongoDB, trading.ID)
	if err != nil {
		return nil, err
	}

	if len(tradingPositions) == 0 {
		return trading, nil
	}

	firstPosition := tradingPositions[len(tradingPositions)-1]
	lastPosition := tradingPositions[0]
	var lastClosedPosition *mongodbModels.Position

	for _, p := range tradingPositions {
		if lastClosedPosition == nil && p.ClosedAt != nil {
			lastClosedPosition = p
			break
		}
	}

	var currentDepositInBaseCurrency float64
	var currentDepositInSecondaryCurrency float64
	var roiInPercent float64
	var roiInBaseCurrency float64

	if lastClosedPosition != nil {
		startBaseCurrencyAmount := firstPosition.Orders[actionsEnum.Buy].SourceCurrencyAmount
		finalLastClosedPositionBaseCurrencyAmount := lastClosedPosition.Orders[actionsEnum.Sell].TargetCurrencyAmount

		roiInPercent = float64((finalLastClosedPositionBaseCurrencyAmount - startBaseCurrencyAmount) / startBaseCurrencyAmount)
		roiInBaseCurrency = float64(finalLastClosedPositionBaseCurrencyAmount - startBaseCurrencyAmount)
	} else {
		roiInPercent = float64(0)
		roiInBaseCurrency = float64(0)
	}

	if lastPosition.ClosedAt == nil {
		currentDepositInBaseCurrency = float64(0)
		currentDepositInSecondaryCurrency = float64(lastPosition.Orders[actionsEnum.Buy].TargetCurrencyAmount)
	} else {
		currentDepositInBaseCurrency = float64(lastPosition.Orders[actionsEnum.Sell].TargetCurrencyAmount)
		currentDepositInSecondaryCurrency = float64(0)
	}

	return ts.UpdateTrading(ctx, mongoDB, graphqlApi.UpdateTradingInput{
		ID:                                id,
		BaseDepositInBaseCurrency:         &trading.BaseDepositInBaseCurrency,
		CurrentDepositInBaseCurrency:      &currentDepositInBaseCurrency,
		CurrentDepositInSecondaryCurrency: &currentDepositInSecondaryCurrency,
		RoiInPercent:                      &roiInPercent,
		RoiInBaseCurrency:                 &roiInBaseCurrency,
	})
}

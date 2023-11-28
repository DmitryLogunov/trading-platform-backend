package gqlServices

import (
	"context"
	"errors"
	graphqlApi "github.com/DmitryLogunov/trading-platform-backend/internal/app/graphql-api"
	"github.com/DmitryLogunov/trading-platform-backend/internal/core/providers/binance-api-client"
)

type ChartsService struct{}

// GetCandlesticksCharts returns candlestick chart data via request to Binance API /api/v3/klines
func (cs *ChartsService) GetCandlesticksCharts(
	ctx context.Context,
	binanceApiClient *binance_api_client.BinanceAPIClient,
	ticker string) ([]*graphqlApi.Candlestick, error) {
	return nil, errors.New("not implemented")
}

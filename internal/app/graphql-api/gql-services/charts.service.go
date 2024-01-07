package gqlServices

import (
	"context"
	graphqlApi "github.com/DmitryLogunov/trading-platform-backend/internal/app/graphql-api"
	"github.com/DmitryLogunov/trading-platform-backend/internal/core/providers/binance-api-client"
	marketTypes "github.com/DmitryLogunov/trading-platform-backend/internal/core/providers/binance-api-client/enums/market-types"
	"github.com/DmitryLogunov/trading-platform-backend/internal/core/providers/binance-api-client/enums/timeframes"
)

type ChartsService struct{}

// GetCandlesticksChart returns candlestick chart data via request to Binance API /api/v3/klines
func (cs *ChartsService) GetCandlesticksChart(
	ctx context.Context,
	binanceApiClient *binance_api_client.BinanceAPIClient,
	ticker string,
) ([]*graphqlApi.Candlestick, error) {
	data, err := binanceApiClient.GetCandlesticksChart(marketTypes.Spot, ticker, timeframes.OneMin)

	if err != nil {
		return nil, err
	}

	return data, nil
}

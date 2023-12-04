package binance_api_client

import (
	marketTypes "github.com/DmitryLogunov/trading-platform-backend/internal/core/providers/binance-api-client/enums/market-types"
	"github.com/DmitryLogunov/trading-platform-backend/internal/core/providers/binance-api-client/enums/timeframes"
	"testing"
)

// TestGetPrices
func TestGetCandlesticksChart(t *testing.T) {
	bp := &BinanceAPIClient{}
	ticker := "BTCUSDT"

	candlesticks, err := bp.GetCandlesticksChart(marketTypes.Spot, ticker, timeframes.OneMin)

	if err != nil {
		t.Fatalf(`GetCandlesticksChart error: %s`, err)
	}

	if candlesticks == nil {
		t.Fatalf(`GetCandlesticksChart error: candlsticks list is empty`)
	}
}

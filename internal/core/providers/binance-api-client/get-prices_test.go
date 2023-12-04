package binance_api_client

import (
	marketTypes "github.com/DmitryLogunov/trading-platform-backend/internal/core/providers/binance-api-client/enums/market-types"
	"testing"
)

// TestGetPrices
func TestGetPrices(t *testing.T) {
	bp := &BinanceAPIClient{}

	tickers := []string{"BTCUSDT", "LTCUSDT"}
	prisesList := bp.GetPrices(marketTypes.Spot, tickers)

	if prisesList == nil {
		t.Fatalf(`GetPrices error: prisesList is nil`)
	}

	if (*prisesList.Data)[0].Symbol != "BTCUSDT" {
		t.Fatalf(`GetPrices error: the symbol of the first item in the prices list is not a BTCUSDT`)
	}

	if (*prisesList.Data)[1].Symbol != "LTCUSDT" {
		t.Fatalf(`GetPrices error: the symbol of the first item in the prices list is not a LTCUSDT`)
	}
}

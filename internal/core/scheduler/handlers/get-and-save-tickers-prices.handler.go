package handlers

import (
	binanceAPIClient "github.com/DmitryLogunov/trading-platform-backend/internal/core/providers/binance-api-client"
	marketTypes "github.com/DmitryLogunov/trading-platform-backend/internal/core/providers/binance-api-client/enums/market-types"
)

type GetTickerPricesParams struct {
	Tickers []string
}

func (hs *Storage) GetAndSaveTickerPricesHandler(c *binanceAPIClient.BinanceAPIClient) func([]HandlerParam) func(interface{}) bool {
	return func(params []HandlerParam) func(interface{}) bool {
		return func(interface{}) bool {
			var tickers []string
			if params != nil && len(params) > 0 {
				tickers = []string{}
				for _, p := range params {
					tickers = append(tickers, p.Value)
				}
			}

			c.GetPrices(marketTypes.Spot, tickers)
			return true
		}
	}
}

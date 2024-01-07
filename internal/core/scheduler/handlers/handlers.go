package handlers

import binanceAPIClient "github.com/DmitryLogunov/trading-platform-backend/internal/core/providers/binance-api-client"

type HandlerParam struct {
	Key   string
	Value string
}

type Storage struct {
	data map[string]func([]HandlerParam) func(interface{}) bool
}

func (hs *Storage) Init() {
	hs.data = make(map[string]func([]HandlerParam) func(interface{}) bool)

	hs.data["do-something-tag"] = hs.DoSomethingHandler

	binanceAPIClient := binanceAPIClient.BinanceAPIClient{}
	hs.data["get-and-save-tickers-prices"] = hs.GetAndSaveTickerPricesHandler(&binanceAPIClient)
}

func (hs *Storage) GetHandler(handlerTag string) func([]HandlerParam) func(interface{}) bool {
	return hs.data[handlerTag]
}

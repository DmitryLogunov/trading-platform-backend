package binance_api_client

import "time"

type Price struct {
	Symbol string
	Price  string
}

type TickersPricesList struct {
	Datetime time.Time
	Data     *[]Price
}

type Candlestick struct {
	Datetime int64
	Data     []float64
}

type BinanceAPIClient struct {
}

type Kline struct {
	OpenTime                 int64
	Open                     string
	High                     string
	Low                      string
	Close                    string
	Volume                   string
	CloseTime                int64
	QuoteAssetVolume         string
	TakerBuyBaseAssetVolume  string
	TakerBuyQuoteAssetVolume string
	Trades                   int64
}

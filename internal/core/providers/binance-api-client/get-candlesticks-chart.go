package binance_api_client

import (
	"context"
	"errors"
	"fmt"
	graphqlApi "github.com/DmitryLogunov/trading-platform-backend/internal/app/graphql-api"
	marketTypes "github.com/DmitryLogunov/trading-platform-backend/internal/core/providers/binance-api-client/enums/market-types"
	"github.com/DmitryLogunov/trading-platform-backend/internal/core/providers/binance-api-client/enums/timeframes"
	"github.com/adshao/go-binance/v2"
	"strconv"
)

var (
	apiKey    = "your api key"
	secretKey = "your secret key"
)

// GetCandlesticksChart returns candlesticks list as a response of Binance API request /api/v3/klines in chart required format
// see: https://apexcharts.com/docs/chart-types/candlestick/
func (bc *BinanceAPIClient) GetCandlesticksChart(marketType uint, ticker string, timeframeTag uint) ([]*graphqlApi.Candlestick, error) {
	timeframe, err := timeframes.GetTimeframe(timeframeTag)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var klines []Kline
	if marketType == marketTypes.Spot {
		klines, err = bc.getKlinesSpot(ticker, timeframe)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	} else if marketType == marketTypes.Futures {
		klines, err = bc.getKlinesFutures(ticker, timeframe)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	} else {
		return nil, errors.New("unknown market type")
	}

	var candlesticks []*graphqlApi.Candlestick
	for _, k := range klines {
		prices, err := bc.parseKlinePrices(&k)
		if err != nil {
			continue
		}

		candlesticks = append(candlesticks, &graphqlApi.Candlestick{
			X: int(k.OpenTime),
			Y: prices,
		})
	}

	return candlesticks, nil
}

// parseKlinePrices: parses Kline prices to []float64 array data
func (bc *BinanceAPIClient) parseKlinePrices(k *Kline) (data []float64, err error) {
	openPrice, err := strconv.ParseFloat(k.Open, 64)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	highPrice, err := strconv.ParseFloat(k.High, 64)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	lowPrice, err := strconv.ParseFloat(k.Low, 64)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	closePrice, err := strconv.ParseFloat(k.Close, 64)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return []float64{openPrice, highPrice, lowPrice, closePrice}, nil
}

func (bc *BinanceAPIClient) getKlinesSpot(ticker string, timeframe string) (klines []Kline, err error) {
	klinesService := binance.NewClient(apiKey, secretKey).NewKlinesService()

	binanceApiKlines, err := klinesService.Symbol(ticker).Interval(timeframe).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for _, k := range binanceApiKlines {
		klines = append(klines, Kline{
			OpenTime:                 k.OpenTime,
			Open:                     k.Open,
			High:                     k.High,
			Low:                      k.Low,
			Close:                    k.Close,
			Volume:                   k.Volume,
			CloseTime:                k.CloseTime,
			QuoteAssetVolume:         k.QuoteAssetVolume,
			TakerBuyBaseAssetVolume:  k.TakerBuyBaseAssetVolume,
			TakerBuyQuoteAssetVolume: k.TakerBuyQuoteAssetVolume,
			Trades:                   k.TradeNum,
		})
	}

	return klines, nil
}

func (bc *BinanceAPIClient) getKlinesFutures(ticker string, timeframe string) (klines []Kline, err error) {
	klinesService := binance.NewFuturesClient(apiKey, secretKey).NewKlinesService()

	binanceApiKlines, err := klinesService.Symbol(ticker).Interval(timeframe).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for _, k := range binanceApiKlines {
		klines = append(klines, Kline{
			OpenTime:                 k.OpenTime,
			Open:                     k.Open,
			High:                     k.High,
			Low:                      k.Low,
			Close:                    k.Close,
			Volume:                   k.Volume,
			CloseTime:                k.CloseTime,
			QuoteAssetVolume:         k.QuoteAssetVolume,
			TakerBuyBaseAssetVolume:  k.TakerBuyBaseAssetVolume,
			TakerBuyQuoteAssetVolume: k.TakerBuyQuoteAssetVolume,
			Trades:                   k.TradeNum,
		})
	}

	return klines, nil
}

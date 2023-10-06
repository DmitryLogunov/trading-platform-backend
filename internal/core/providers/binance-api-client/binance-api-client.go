package binance_api_client

import (
	"encoding/json"
	"fmt"
	marketTypes "github.com/DmitryLogunov/trading-platform-backend/internal/core/providers/binance-api-client/enums/market-types"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Price struct {
	Symbol string
	Price  string
}

type TickersPricesList struct {
	Datetime time.Time
	Data     *[]Price
}

const BinanceSpotApiUrl = "https://api.binance.com"
const BinanceFuturesApiUrl = "https://fapi.binance.com"

type BinanceAPIClient struct {
}

func (bc *BinanceAPIClient) GetPrices(marketType uint, tickers []string) *TickersPricesList {
	var baseAPIUrl string
	if marketType == marketTypes.Spot {
		baseAPIUrl = BinanceSpotApiUrl
	} else if marketType == marketTypes.Futures {
		baseAPIUrl = BinanceFuturesApiUrl
	} else {
		fmt.Println("Error: unknown market type. API url is undefined")
		return nil
	}

	tickersFilter := ""
	if tickers != nil && len(tickers) > 0 {
		tickersFilter = fmt.Sprintf("?symbols=[\"%s\"]", strings.Join(tickers, "\",\""))
	}

	url := fmt.Sprintf("%s/api/v3/ticker/price%s", baseAPIUrl, tickersFilter)
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		return nil
	}

	var tickersPricesList []Price
	err = json.NewDecoder(res.Body).Decode(&tickersPricesList)
	if err != nil {
		fmt.Printf("client: could not parse response body: %s\n", err)
		return nil
	}

	for _, tickerPrice := range tickersPricesList {
		last4chars := tickerPrice.Symbol[len(tickerPrice.Symbol)-4:]
		if last4chars != "USDT" {
			continue
		}

		price, err := strconv.ParseFloat(tickerPrice.Price, 32)
		if err != nil {
			fmt.Printf("error: %s, could not parse price: %s\n", err, tickerPrice.Price)
			return nil
		}

		fmt.Printf("symbol: %s, price: %f\n", tickerPrice.Symbol, float32(price))
	}

	return nil
}

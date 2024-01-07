package binance_api_client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// GetPrices returns prices list as a response of Binance API request /api/v3/ticker/price
func (bc *BinanceAPIClient) GetPrices(marketType uint, tickers []string) *TickersPricesList {
	baseAPIUrl, err := bc.getBinanceApiBaseUrl(marketType)
	if err != nil {
		fmt.Printf("client: could get Binance base API Url: %s\n", err)
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

	return &TickersPricesList{
		Datetime: time.Now(),
		Data:     &tickersPricesList,
	}
}

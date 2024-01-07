package binance_api_client

import (
	"errors"
	marketTypes "github.com/DmitryLogunov/trading-platform-backend/internal/core/providers/binance-api-client/enums/market-types"
)

func (bc *BinanceAPIClient) getBinanceApiBaseUrl(marketType uint) (string, error) {
	if marketType == marketTypes.Spot {
		return BinanceSpotApiUrl, nil
	} else if marketType == marketTypes.Futures {
		return BinanceFuturesApiUrl, nil
	} else {
		return "", errors.New("unknown market type. API url is undefined")
	}
}

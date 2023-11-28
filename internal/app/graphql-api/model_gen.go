// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphql_api

import (
	"time"
)

type Alert struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Ticker    string    `json:"ticker"`
	Action    int       `json:"action"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"createdAt"`
}

type AlertsFiltersInput struct {
	Title         *string `json:"title,omitempty"`
	Ticker        *string `json:"ticker,omitempty"`
	Action        *string `json:"action,omitempty"`
	CreatedAtFrom *string `json:"createdAtFrom,omitempty"`
	CreatedAtTo   *string `json:"createdAtTo,omitempty"`
}

type Candlestick struct {
	Datetime time.Time `json:"datetime"`
	Data     []float64 `json:"data"`
}

type CandlesticksChartFiltersInput struct {
	Ticker string `json:"ticker"`
}

type ClosePositionInput struct {
	ID       string  `json:"id"`
	Price    float64 `json:"price"`
	ClosedAt *string `json:"closedAt,omitempty"`
}

type CronPeriodInput struct {
	Unit     string `json:"unit"`
	Interval int    `json:"interval"`
}

type CronPeriodObject struct {
	Unit     string `json:"unit"`
	Interval int    `json:"interval"`
}

type Job struct {
	Tag        string    `json:"tag"`
	HandlerTag string    `json:"handlerTag"`
	Params     string    `json:"params"`
	CronPeriod string    `json:"cronPeriod"`
	CreatedAt  time.Time `json:"createdAt"`
	Status     string    `json:"status"`
}

type JobData struct {
	HandlerTag string           `json:"handlerTag"`
	Params     []*JobParamInput `json:"params"`
	CronPeriod *CronPeriodInput `json:"cronPeriod"`
}

type JobParamInput struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type JobParamObject struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}

type NewTradingInput struct {
	Exchange                  string  `json:"exchange"`
	BaseCurrency              string  `json:"baseCurrency"`
	SecondaryCurrency         string  `json:"secondaryCurrency"`
	BaseDepositInBaseCurrency float64 `json:"baseDepositInBaseCurrency"`
	StartedAt                 *string `json:"startedAt,omitempty"`
}

type OpenPositionInput struct {
	TradingID          string  `json:"tradingId"`
	BaseCurrencyAmount float64 `json:"baseCurrencyAmount"`
	Price              float64 `json:"price"`
	CreatedAt          *string `json:"createdAt,omitempty"`
}

type Order struct {
	Action               int       `json:"action"`
	SourceCurrencyAmount float64   `json:"sourceCurrencyAmount"`
	TargetCurrencyAmount float64   `json:"targetCurrencyAmount"`
	Price                float64   `json:"price"`
	CreatedAt            time.Time `json:"createdAt"`
}

type Position struct {
	ID                string     `json:"id"`
	TradingID         string     `json:"tradingId"`
	BaseCurrency      string     `json:"baseCurrency"`
	SecondaryCurrency string     `json:"secondaryCurrency"`
	Orders            []*Order   `json:"orders"`
	RoiInPercent      *float64   `json:"roiInPercent,omitempty"`
	RoiInBaseCurrency *float64   `json:"roiInBaseCurrency,omitempty"`
	CreatedAt         time.Time  `json:"createdAt"`
	ClosedAt          *time.Time `json:"closedAt,omitempty"`
}

type Price struct {
	Ticker   string    `json:"ticker"`
	Price    float64   `json:"price"`
	Datetime time.Time `json:"datetime"`
}

type Trading struct {
	ID                                string     `json:"id"`
	Exchange                          string     `json:"exchange"`
	BaseCurrency                      string     `json:"baseCurrency"`
	SecondaryCurrency                 string     `json:"secondaryCurrency"`
	BaseDepositInBaseCurrency         float64    `json:"baseDepositInBaseCurrency"`
	CurrentDepositInBaseCurrency      *float64   `json:"currentDepositInBaseCurrency,omitempty"`
	CurrentDepositInSecondaryCurrency *float64   `json:"currentDepositInSecondaryCurrency,omitempty"`
	RoiInPercent                      *float64   `json:"roiInPercent,omitempty"`
	RoiInBaseCurrency                 *float64   `json:"roiInBaseCurrency,omitempty"`
	StartedAt                         time.Time  `json:"startedAt"`
	ClosedAt                          *time.Time `json:"closedAt,omitempty"`
}

type UpdateTradingInput struct {
	ID                                string   `json:"id"`
	BaseDepositInBaseCurrency         *float64 `json:"baseDepositInBaseCurrency,omitempty"`
	CurrentDepositInBaseCurrency      *float64 `json:"currentDepositInBaseCurrency,omitempty"`
	CurrentDepositInSecondaryCurrency *float64 `json:"currentDepositInSecondaryCurrency,omitempty"`
	RoiInPercent                      *float64 `json:"roiInPercent,omitempty"`
	RoiInBaseCurrency                 *float64 `json:"roiInBaseCurrency,omitempty"`
	ClosedAt                          *string  `json:"closedAt,omitempty"`
}

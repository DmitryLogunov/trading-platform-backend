package actions

import (
	"errors"
	"strings"
)

const (
	Buy = iota
	Sell
	ErrorActionValue
)

func Parse(str string) (uint, error) {
	if strings.ToLower(str) == "buy" {
		return Buy, nil
	}

	if strings.ToLower(str) == "sell" {
		return Sell, nil
	}

	return ErrorActionValue, errors.New("wrong alert action format")
}

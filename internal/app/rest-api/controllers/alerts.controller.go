package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/DmitryLogunov/trading-platform/internal/app/rest-api/dto"
	alertActions "github.com/DmitryLogunov/trading-platform/internal/core/database/mongodb/enums/alert-actions"
	mongodbModels "github.com/DmitryLogunov/trading-platform/internal/core/database/mongodb/models"
	"github.com/DmitryLogunov/trading-platform/internal/core/helpers"
	"net/http"
	"strconv"
)

func (c *Controllers) AddAlert(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	alertData := &dto.AlertDTO{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(alertData)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	createdAt, err := helpers.DatetimeParse(alertData.CreatedAt)
	if err != nil {
		http.Error(w, "Wrong datetime format of createdAt", http.StatusBadRequest)
		return
	}

	action, err := alertActions.Parse(alertData.Action)
	if err != nil {
		http.Error(w, "Wrong alert action format", http.StatusBadRequest)
		return
	}

	price, err := strconv.ParseFloat(alertData.Price, 32)
	if err != nil {
		http.Error(w, "Wrong alert action format", http.StatusBadRequest)
		return
	}

	alert := &mongodbModels.Alert{
		Title:     alertData.Title,
		Ticker:    alertData.Ticker,
		Action:    action,
		Price:     float32(price),
		CreatedAt: *createdAt,
	}

	alert.Save(r.Context(), c.MongoDB, alert)

	fmt.Printf("New alert saved: {title: %s, ticker: %s, action: %d, price: %0.6f, createdAt: %s}\n", alert.Title, alert.Ticker, alert.Action, alert.Price, alert.CreatedAt)

	w.WriteHeader(http.StatusCreated)
}

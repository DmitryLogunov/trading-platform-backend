package dto

type AlertDTO struct {
	Title     string `json:"title",gorm:"not null"`
	Ticker    string `json:"ticker",gorm:"not null"`
	Action    string `json:"action",gorm:"not null"`
	Price     string `json:"price",gorm:"not null"`
	CreatedAt string `json:"createdAt",gorm:"not null"`
}

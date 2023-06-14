package dbmodel

import "time"

type Post struct {
	ID          uint      `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Title       string    `gorm:"not null"`
	Content     string    `gorm:"not null"`
	Author      string    `gorm:"not null; unique"`
	Hero        string    `json:"Hero"`
	PublishedAt time.Time `json:"PublishedAt"`
	UpdatedAt   time.Time `json:"UpdateAt"`
}

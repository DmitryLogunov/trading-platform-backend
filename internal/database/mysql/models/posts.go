package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type Post struct {
	Title       string    `gorm:"not null"`
	Content     string    `gorm:"not null"`
	Author      string    `gorm:"not null; unique"`
	Hero        string    `json:"Hero"`
	PublishedAt time.Time `json:"PublishedAt,omitempty"`
	UpdatedAt   time.Time `json:"UpdateAt,omitempty"`
}

// CreatePost : creates posts
func (p *Post) CreatePost(input *Post, db *gorm.DB) (*Post, error) {
	addedPost := Post{
		Title:       input.Title,
		Content:     input.Content,
		Author:      input.Author,
		Hero:        input.Hero,
		PublishedAt: time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := db.Create(&addedPost).Error; err != nil {
		fmt.Println(err)
		return nil, err

	}

	return &addedPost, nil
}

// GetPosts : returns posts
func (p *Post) GetPosts(db *gorm.DB) ([]*Post, error) {
	var posts []*Post

	GetPosts := db.Model(&posts).Find(&posts)

	if GetPosts.Error != nil {
		fmt.Println(GetPosts.Error)
		return nil, GetPosts.Error
	}

	return posts, nil
}

package models

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Post struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title",gorm:"not null"`
	Content     string             `bson:"content",gorm:"not null"`
	Author      string             `bson:"author",gorm:"not null; unique"`
	Hero        string             `bson:"hero"`
	PublishedAt time.Time          `bson:"created_at,omitempty"`
	UpdatedAt   time.Time          `bson:"updated_at,omitempty"`
}

// getCollection: returns "posts" mongodb collection
func (p *Post) getCollection(db *mongo.Database) *mongo.Collection {
	collectionName := "posts"

	return db.Collection(collectionName)
}

// CreatePost : creates posts
func (p *Post) CreatePost(ctx context.Context, input *Post, db *mongo.Database) (*Post, error) {
	addedPost := Post{
		Title:       input.Title,
		Content:     input.Content,
		Author:      input.Author,
		Hero:        input.Hero,
		PublishedAt: time.Now(),
		UpdatedAt:   time.Now(),
	}

	res, err := p.getCollection(db).InsertOne(ctx, addedPost)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	addedPost.ID = res.InsertedID.(primitive.ObjectID)

	return &addedPost, nil
}

// GetPosts : returns posts
func (p *Post) GetPosts(ctx context.Context, db *mongo.Database) ([]*Post, error) {
	cursor, err := p.getCollection(db).Find(ctx, bson.M{})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var posts []*Post
	for cursor.Next(context.Background()) {
		post := Post{}
		cursor.Decode(&post)

		posts = append(posts, &post)
	}

	return posts, nil
}

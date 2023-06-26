package gqlServices

import (
	"context"
	"fmt"
	graphqlApi "github.com/DmitryLogunov/trading-platform/internal/app/graphql-api"
	mongodbModels "github.com/DmitryLogunov/trading-platform/internal/core/database/mongodb/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostsService struct{}

func (ps *PostsService) CreatePost(ctx context.Context, mongoDB *mongo.Database, input graphqlApi.NewPost) (*graphqlApi.Post, error) {
	postsModel := mongodbModels.Post{
		Title:   input.Title,
		Content: input.Content,
		Author:  input.Author,
		Hero:    input.Hero,
	}

	addedPost, err := postsModel.CreatePost(ctx, &postsModel, mongoDB)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &graphqlApi.Post{
		ID:          addedPost.ID.Hex(),
		Title:       addedPost.Title,
		Content:     addedPost.Content,
		Author:      addedPost.Author,
		Hero:        addedPost.Hero,
		PublishedAt: addedPost.PublishedAt,
		UpdatedAt:   addedPost.UpdatedAt,
	}, nil
}

func (ps *PostsService) GetPosts(ctx context.Context, mongoDB *mongo.Database) ([]*graphqlApi.Post, error) {
	postsModel := mongodbModels.Post{}

	posts, err := postsModel.GetPosts(ctx, mongoDB)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var gqlPosts []*graphqlApi.Post

	for _, p := range posts {
		gqlPosts = append(gqlPosts, &graphqlApi.Post{
			ID:          p.ID.Hex(),
			Title:       p.Title,
			Content:     p.Content,
			Author:      p.Author,
			Hero:        p.Hero,
			PublishedAt: p.PublishedAt,
			UpdatedAt:   p.UpdatedAt,
		})
	}

	return gqlPosts, nil
}

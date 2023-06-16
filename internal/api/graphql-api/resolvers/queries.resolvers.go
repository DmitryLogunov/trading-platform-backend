package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.33

import (
	"context"
	"fmt"

	graphql_api "github.com/DmitryLogunov/trading-platform/internal/api/graphql-api"
	mongodbModels "github.com/DmitryLogunov/trading-platform/internal/database/mongodb/models"
)

// GetPosts is the resolver for the getPosts field.
func (r *queryResolver) GetPosts(ctx context.Context) ([]*graphql_api.Post, error) {
	postsModel := mongodbModels.Post{}

	posts, err := postsModel.GetPosts(ctx, r.MongoDB)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var gqlPosts []*graphql_api.Post

	for _, p := range posts {
		gqlPosts = append(gqlPosts, &graphql_api.Post{
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

// Query returns graphql_api.QueryResolver implementation.
func (r *Resolver) Query() graphql_api.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

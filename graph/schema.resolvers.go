package graph

import (
	"context"
	model "echo-graphql-jwt/models/graph"
)

func (r *mutationResolver) CreatePost(ctx context.Context, input model.CreatePostInput) (*model.Post, error) {
	post := model.Post{
		ID:      input.ID,
		Title:   input.Title,
		Content: input.Content,
	}

	if err := r.DB.Create(&post).Error; err != nil {
		return nil, err
	}

	return &post, nil
}

func (r *mutationResolver) UpdatePost(ctx context.Context, input model.UpdatePostInput) (*model.Post, error) {
	post := model.Post{}

	if err := r.DB.First(&post, input.ID).Error; err != nil {
		return nil, err
	}

	post.Title = *input.Title
	post.Content = *input.Content

	if err := r.DB.Updates(&post).Error; err != nil {
		return nil, err
	}

	return &post, nil
}

func (r *mutationResolver) DeletePost(ctx context.Context, input model.DeletePostInput) (*model.Post, error) {
	post := model.Post{}

	if err := r.DB.First(&post, input.ID).Error; err != nil {
		return nil, err
	}

	if err := r.DB.Delete(&post).Error; err != nil {
		return nil, err
	}

	return &post, nil
}

func (r *queryResolver) Post(ctx context.Context, id string) (*model.Post, error) {
	post := model.Post{}

	if err := r.DB.First(&post, id).Error; err != nil {
		return nil, err
	}

	return &post, nil
}

func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	posts := []*model.Post{}

	if err := r.DB.Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

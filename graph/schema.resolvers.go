package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"algorithm-1/graph/generated"
	"algorithm-1/graph/model"
	"context"
)

func (r *mutationResolver) GenerateSchedule(ctx context.Context, courses []*model.NewCourse, professors []*model.NewProfessor, numSchedules *int) (*model.Schedule, error) {
	// TODO: Trigger schedule generation from this end point.
	return nil, nil
}

func (r *queryResolver) Schedules(ctx context.Context) ([]*model.Schedule, error) {
	// TODO: Return generated schedules as an array
	return nil, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

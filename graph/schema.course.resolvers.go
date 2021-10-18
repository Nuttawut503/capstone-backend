package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Nuttawut503/capstone-backend/graph/generated"
	"github.com/Nuttawut503/capstone-backend/graph/model"
)

func (r *mutationResolver) CreateCourse(ctx context.Context, input model.CreateCourseInput) (*model.Course, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateLOs(ctx context.Context, input model.CreateLOSetInput) ([]string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Courses(ctx context.Context, programID string) ([]*model.Course, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Course(ctx context.Context, courseID string) (*model.Course, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Los(ctx context.Context, courseID string) ([]*model.Lo, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

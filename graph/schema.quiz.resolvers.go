package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Nuttawut503/capstone-backend/graph/model"
)

func (r *mutationResolver) CreateQuiz(ctx context.Context, input *model.CreateQuizInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateQuestionLink(ctx context.Context, input *model.CreateQuestionLinkInput) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Quiz(ctx context.Context, courseID string) ([]*model.Quiz, error) {
	panic(fmt.Errorf("not implemented"))
}

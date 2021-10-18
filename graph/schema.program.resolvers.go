package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Nuttawut503/capstone-backend/graph/model"
)

func (r *mutationResolver) CreateProgram(ctx context.Context, input model.CreateProgramInput) (*model.Program, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreatePLOGroupInput(ctx context.Context, input model.CreatePLOGroupInput) (*model.PLOGroup, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Programs(ctx context.Context) ([]*model.Program, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) PloGroups(ctx context.Context, programID string) ([]*model.PLOGroup, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Plos(ctx context.Context, ploGroupID string) ([]*model.Plo, error) {
	panic(fmt.Errorf("not implemented"))
}

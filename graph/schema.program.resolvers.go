package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Nuttawut503/capstone-backend/db"
	"github.com/Nuttawut503/capstone-backend/graph/model"
	"github.com/prisma/prisma-client-go/runtime/transaction"
)

func (r *mutationResolver) CreateProgram(ctx context.Context, input model.CreateProgramInput) (*model.Program, error) {
	createdProgram, err := r.Client.Program.CreateOne(
		db.Program.Name.Set(input.Name),
		db.Program.Description.Set(input.Description),
	).Exec(ctx)
	if err != nil {
		return &model.Program{}, err
	}
	return &model.Program{
		ID:          createdProgram.ID,
		Name:        createdProgram.Name,
		Description: createdProgram.Description,
	}, nil
}

func (r *mutationResolver) CreatePLOGroup(ctx context.Context, programID string, name string, input []*model.CreatePLOsInput) (*model.PLOGroup, error) {
	createPLOGroup, err := r.Client.PLOgroup.CreateOne(
		db.PLOgroup.Name.Set(name),
		db.PLOgroup.Program.Link(
			db.Program.ID.Equals(programID),
		),
	).Exec(ctx)
	if err != nil {
		return &model.PLOGroup{}, err
	}
	transactions := []transaction.Param{}
	for _, plo := range input {
		transactions = append(transactions, r.Client.PLO.CreateOne(
			db.PLO.Title.Set(plo.Title),
			db.PLO.Description.Set(plo.Description),
			db.PLO.PloGroup.Link(
				db.PLOgroup.ID.Equals(createPLOGroup.ID),
			),
		).Tx())
	}
	if err := r.Client.Prisma.Transaction(transactions...).Exec(ctx); err != nil {
		return &model.PLOGroup{}, err
	}
	return &model.PLOGroup{
		ID:   createPLOGroup.ID,
		Name: createPLOGroup.Name,
	}, nil
}

func (r *mutationResolver) CreatePlo(ctx context.Context, ploGroupID string, input model.CreatePLOInput) (*model.Plo, error) {
	createdPLO, err := r.Client.PLO.CreateOne(
		db.PLO.Title.Set(input.Title),
		db.PLO.Description.Set(input.Description),
		db.PLO.PloGroup.Link(
			db.PLOgroup.ID.Equals(ploGroupID),
		),
	).Exec(ctx)
	if err != nil {
		return &model.Plo{}, err
	}
	return &model.Plo{
		ID:          createdPLO.ID,
		Title:       createdPLO.Title,
		Description: createdPLO.Description,
		PloGroupID:  createdPLO.PloGroupID,
	}, nil
}

func (r *mutationResolver) DeletePLOGroup(ctx context.Context, id string) (*model.DeletePLOGroupResult, error) {
	deleted, err := r.Client.PLOgroup.FindUnique(
		db.PLOgroup.ID.Equals(id),
	).Delete().Exec(ctx)
	if err != nil {
		return &model.DeletePLOGroupResult{}, err
	}
	return &model.DeletePLOGroupResult{
		ID: deleted.ID,
	}, nil
}

func (r *mutationResolver) DeletePlo(ctx context.Context, id string) (*model.DeletePLOResult, error) {
	deleted, err := r.Client.PLO.FindUnique(
		db.PLO.ID.Equals(id),
	).Delete().Exec(ctx)
	if err != nil {
		return &model.DeletePLOResult{}, err
	}
	return &model.DeletePLOResult{
		ID: deleted.ID,
	}, nil
}

func (r *queryResolver) Programs(ctx context.Context) ([]*model.Program, error) {
	allPrograms, err := r.Client.Program.FindMany().Exec(ctx)
	if err != nil {
		return []*model.Program{}, err
	}
	programs := []*model.Program{}
	for _, program := range allPrograms {
		programs = append(programs, &model.Program{
			ID:          program.ID,
			Name:        program.Name,
			Description: program.Description,
		})
	}
	return programs, nil
}

func (r *queryResolver) Program(ctx context.Context, programID string) (*model.Program, error) {
	program, err := r.Client.Program.FindUnique(
		db.Program.ID.Equals(programID),
	).Exec(ctx)
	if err != nil {
		return &model.Program{}, err
	}
	return &model.Program{
		ID:          program.ID,
		Name:        program.Name,
		Description: program.Description,
	}, nil
}

func (r *queryResolver) PloGroups(ctx context.Context, programID string) ([]*model.PLOGroup, error) {
	allPLOGroups, err := r.Client.PLOgroup.FindMany(
		db.PLOgroup.ProgramID.Equals(programID),
	).Exec(ctx)
	if err != nil {
		return []*model.PLOGroup{}, err
	}
	ploGroups := []*model.PLOGroup{}
	for _, ploGroup := range allPLOGroups {
		ploGroups = append(ploGroups, &model.PLOGroup{
			ID:   ploGroup.ID,
			Name: ploGroup.Name,
		})
	}
	return ploGroups, nil
}

func (r *queryResolver) Plos(ctx context.Context, ploGroupID string) ([]*model.Plo, error) {
	allPLOs, err := r.Client.PLO.FindMany(
		db.PLO.PloGroupID.Equals(ploGroupID),
	).Exec(ctx)
	if err != nil {
		return []*model.Plo{}, err
	}
	plos := []*model.Plo{}
	for _, plo := range allPLOs {
		plos = append(plos, &model.Plo{
			ID:          plo.ID,
			Title:       plo.Title,
			Description: plo.Description,
		})
	}
	return plos, nil
}

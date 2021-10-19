package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Nuttawut503/capstone-backend/db"
	"github.com/Nuttawut503/capstone-backend/graph/generated"
	"github.com/Nuttawut503/capstone-backend/graph/model"
)

func (r *mutationResolver) CreateCourse(ctx context.Context, programID string, input model.CreateCourseInput) (*model.Course, error) {
	createdCourse, err := r.Client.Course.CreateOne(
		db.Course.Name.Set(input.Name),
		db.Course.Description.Set(input.Description),
		db.Course.Semester.Set(input.Semester),
		db.Course.Year.Set(input.Year),
		db.Course.Program.Link(
			db.Program.ID.Equals(programID),
		),
		db.Course.PloGroup.Link(
			db.PLOgroup.ID.Equals(input.PloGroupID),
		),
	).Exec(ctx)
	if err != nil {
		return &model.Course{}, err
	}
	return &model.Course{
		ID:          createdCourse.ID,
		Name:        createdCourse.Name,
		Description: createdCourse.Description,
		Semester:    createdCourse.Semester,
		Year:        createdCourse.Year,
		PloGroupID:  createdCourse.PloGroupID,
	}, nil
}

func (r *mutationResolver) CreateLOs(ctx context.Context, courseID string, input []*model.CreateLOsInput) ([]*model.CreateLOResult, error) {
	result := []*model.CreateLOResult{}
	for _, loInput := range input {
		createdLO, err := r.Client.LO.CreateOne(
			db.LO.Title.Set(loInput.Title),
			db.LO.Course.Link(
				db.Course.ID.Equals(courseID),
			),
		).Exec(ctx)
		if err != nil {
			return []*model.CreateLOResult{}, err
		}
		for _, loLevelInput := range loInput.Levels {
			r.CreateLOLevel(ctx, createdLO.ID, *loLevelInput)
		}
		result = append(result, &model.CreateLOResult{
			ID: createdLO.ID,
		})
	}
	return result, nil
}

func (r *mutationResolver) CreateLOLink(ctx context.Context, loID string, ploID string) (*model.CreateLOLinkResult, error) {
	createdLO, err := r.Client.LOlink.CreateOne(
		db.LOlink.Lo.Link(
			db.LO.ID.Equals(loID),
		),
		db.LOlink.Plo.Link(
			db.PLO.ID.Equals(ploID),
		),
	).Exec(ctx)
	if err != nil {
		return &model.CreateLOLinkResult{}, err
	}
	return &model.CreateLOLinkResult{
		LoID:  createdLO.LoID,
		PloID: createdLO.PloID,
	}, nil
}

func (r *mutationResolver) CreateLo(ctx context.Context, courseID string, input model.CreateLOInput) (*model.CreateLOResult, error) {
	createdLO, err := r.Client.LO.CreateOne(
		db.LO.Title.Set(input.Title),
		db.LO.Course.Link(
			db.Course.ID.Equals(courseID),
		),
	).Exec(ctx)
	if err != nil {
		return &model.CreateLOResult{}, err
	}
	_, err = r.Client.LOlevel.CreateOne(
		db.LOlevel.Level.Set(input.Level),
		db.LOlevel.Description.Set(input.Description),
		db.LOlevel.Lo.Link(
			db.LO.ID.Equals(createdLO.ID),
		),
	).Exec(ctx)
	if err != nil {
		return &model.CreateLOResult{}, err
	}
	return &model.CreateLOResult{
		ID: createdLO.ID,
	}, nil
}

func (r *mutationResolver) CreateLOLevel(ctx context.Context, loID string, input model.CreateLOLevelInput) (*model.CreateLOResult, error) {
	createdLOLevel, err := r.Client.LOlevel.CreateOne(
		db.LOlevel.Level.Set(input.Level),
		db.LOlevel.Description.Set(input.Description),
		db.LOlevel.Lo.Link(
			db.LO.ID.Equals(loID),
		),
	).Exec(ctx)
	if err != nil {
		return &model.CreateLOResult{}, err
	}
	return &model.CreateLOResult{
		ID: createdLOLevel.LoID,
	}, nil
}

func (r *mutationResolver) DeleteLo(ctx context.Context, id string) (*model.DeleteLOResult, error) {
	deleted, err := r.Client.LO.FindUnique(
		db.LO.ID.Equals(id),
	).Delete().Exec(ctx)
	if err != nil {
		return &model.DeleteLOResult{}, err
	}
	return &model.DeleteLOResult{
		ID: deleted.ID,
	}, nil
}

func (r *mutationResolver) DeleteLOLevel(ctx context.Context, id string, level int) (*model.DeleteLOLevelResult, error) {
	deleted, err := r.Client.LOlevel.FindUnique(
		db.LOlevel.LoIDLevel(
			db.LOlevel.LoID.Equals(id),
			db.LOlevel.Level.Equals(level),
		),
	).Delete().Exec(ctx)
	if err != nil {
		return &model.DeleteLOLevelResult{}, err
	}
	return &model.DeleteLOLevelResult{
		ID: deleted.LoID,
	}, nil
}

func (r *mutationResolver) DeleteLOLink(ctx context.Context, loID string, ploID string) (*model.DeleteLOLinkResult, error) {
	deleted, err := r.Client.LOlink.FindUnique(
		db.LOlink.LoIDPloID(
			db.LOlink.LoID.Equals(loID),
			db.LOlink.PloID.Equals(ploID),
		),
	).Delete().Exec(ctx)
	if err != nil {
		return &model.DeleteLOLinkResult{}, err
	}
	return &model.DeleteLOLinkResult{
		LoID:  deleted.LoID,
		PloID: deleted.PloID,
	}, nil
}

func (r *queryResolver) Courses(ctx context.Context, programID string) ([]*model.Course, error) {
	allCourses, err := r.Client.Course.FindMany(
		db.Course.Program.Where(
			db.Program.ID.Equals(programID),
		),
	).Exec(ctx)
	if err != nil {
		return []*model.Course{}, err
	}
	courses := []*model.Course{}
	for _, course := range allCourses {
		courses = append(courses, &model.Course{
			ID:          course.ID,
			Name:        course.Name,
			Description: course.Description,
			Semester:    course.Semester,
			Year:        course.Year,
			PloGroupID:  course.PloGroupID,
		})
	}
	return courses, nil
}

func (r *queryResolver) Course(ctx context.Context, courseID string) (*model.Course, error) {
	course, err := r.Client.Course.FindUnique(
		db.Course.ID.Equals(courseID),
	).Exec(ctx)
	if err != nil {
		return &model.Course{}, err
	}
	return &model.Course{
		ID:          course.ID,
		Name:        course.Name,
		Description: course.Description,
		Semester:    course.Semester,
		Year:        course.Year,
		PloGroupID:  course.PloGroupID,
	}, nil
}

func (r *queryResolver) Los(ctx context.Context, courseID string) ([]*model.Lo, error) {
	allLOs, err := r.Client.LO.FindMany(
		db.LO.Course.Where(
			db.Course.ID.Equals(courseID),
		),
	).With(
		db.LO.Levels.Fetch(),
		db.LO.Links.Fetch().With(
			db.LOlink.Plo.Fetch(),
		),
	).Exec(ctx)
	if err != nil {
		return []*model.Lo{}, err
	}
	los := []*model.Lo{}
	for _, lo := range allLOs {
		levels := []*model.LOLevel{}
		for _, level := range lo.Levels() {
			levels = append(levels, &model.LOLevel{
				Level:       level.Level,
				Description: level.Description,
			})
		}
		ploLinks := []*model.Plo{}
		for _, plo := range lo.Links() {
			ploLinks = append(ploLinks, &model.Plo{
				ID:          plo.PloID,
				Title:       plo.Plo().Title,
				Description: plo.Plo().Description,
				PloGroupID:  plo.Plo().PloGroupID,
			})
		}
		los = append(los, &model.Lo{
			ID:       lo.ID,
			Title:    lo.Title,
			Levels:   levels,
			PloLinks: ploLinks,
		})
	}
	return los, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

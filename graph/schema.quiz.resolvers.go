package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Nuttawut503/capstone-backend/db"
	"github.com/Nuttawut503/capstone-backend/graph/model"
)

func (r *mutationResolver) CreateQuiz(ctx context.Context, courseID string, input *model.CreateQuizInput) (*model.CreateQuizResult, error) {
	createdQuiz, err := r.Client.Quiz.CreateOne(
		db.Quiz.Name.Set(input.Name),
		db.Quiz.Course.Link(
			db.Course.ID.Equals(courseID),
		),
		db.Quiz.CreatedAt.Set(input.CreatedAt),
	).Exec(ctx)
	if err != nil {
		return &model.CreateQuizResult{}, err
	}
	for _, questionInput := range input.Questions {
		createdQuestion, err := r.Client.Question.CreateOne(
			db.Question.Title.Set(questionInput.Title),
			db.Question.MaxScore.Set(questionInput.MaxScore),
			db.Question.Quiz.Link(
				db.Quiz.ID.Equals(createdQuiz.ID),
			),
		).Exec(ctx)
		if err != nil {
			return &model.CreateQuizResult{}, err
		}
		for _, resultInput := range questionInput.Results {
			_, err := r.Client.QuestionResult.CreateOne(
				db.QuestionResult.Question.Link(
					db.Question.ID.Equals(createdQuestion.ID),
				),
				db.QuestionResult.Student.Link(
					db.Student.ID.Equals(resultInput.StudentID),
				),
				db.QuestionResult.Score.Set(resultInput.Score),
			).Exec(ctx)
			if err != nil {
				return &model.CreateQuizResult{}, err
			}
		}
	}
	return &model.CreateQuizResult{
		ID: createdQuiz.ID,
	}, nil
}

func (r *mutationResolver) CreateQuestionLink(ctx context.Context, input *model.CreateQuestionLinkInput) (*model.CreateQuestionLinkResult, error) {
	createdQuestionLink, err := r.Client.QuestionLink.CreateOne(
		db.QuestionLink.Question.Link(
			db.Question.ID.Equals(input.QuestionID),
		),
		db.QuestionLink.LoLevel.Link(
			db.LOlevel.LoIDLevel(
				db.LOlevel.LoID.Equals(input.LoID),
				db.LOlevel.Level.Equals(input.Level),
			),
		),
	).Exec(ctx)
	if err != nil {
		return &model.CreateQuestionLinkResult{}, err
	}
	return &model.CreateQuestionLinkResult{
		QuestionID: createdQuestionLink.QuestionID,
		LoID:       createdQuestionLink.LoID,
	}, nil
}

func (r *mutationResolver) DeleteQuiz(ctx context.Context, id string) (*model.DeleteQuizResult, error) {
	deleted, err := r.Client.Quiz.FindUnique(
		db.Quiz.ID.Equals(id),
	).Delete().Exec(ctx)
	if err != nil {
		return &model.DeleteQuizResult{}, err
	}
	return &model.DeleteQuizResult{
		ID: deleted.ID,
	}, nil
}

func (r *mutationResolver) DeleteQuestionLink(ctx context.Context, input model.DeleteQuestionLinkInput) (*model.DeleteQuestionLinkResult, error) {
	deleted, err := r.Client.QuestionLink.FindUnique(
		db.QuestionLink.QuestionIDLoIDLevel(
			db.QuestionLink.QuestionID.Equals(input.QuestionID),
			db.QuestionLink.LoID.Equals(input.LoID),
			db.QuestionLink.Level.Equals(input.Level),
		),
	).Delete().Exec(ctx)
	if err != nil {
		return &model.DeleteQuestionLinkResult{}, err
	}
	return &model.DeleteQuestionLinkResult{
		QuestionID: deleted.QuestionID,
		LoID:       deleted.LoID,
	}, nil
}

func (r *queryResolver) Quiz(ctx context.Context, courseID string) ([]*model.Quiz, error) {
	quizzes := []*model.Quiz{}
	allQuizzes, err := r.Client.Quiz.FindMany(
		db.Quiz.Course.Where(
			db.Course.ID.Equals(courseID),
		),
	).With(
		db.Quiz.Questions.Fetch().With(
			db.Question.Links.Fetch(),
			db.Question.Results.Fetch(),
		),
	).Exec(ctx)
	if err != nil {
		return []*model.Quiz{}, err
	}
	for _, quiz := range allQuizzes {
		questions := []*model.Question{}
		for _, question := range quiz.Questions() {
			results := []*model.QuestionResult{}
			loLinks := []*model.QuestionLink{}
			for _, result := range question.Results() {
				results = append(results, &model.QuestionResult{
					StudentID: result.StudentID,
					Score:     result.Score,
				})
			}
			for _, loLink := range question.Links() {
				loLinks = append(loLinks, &model.QuestionLink{
					LoID:  loLink.LoID,
					Level: loLink.Level,
				})
			}
			questions = append(questions, &model.Question{
				ID:       question.ID,
				Title:    question.Title,
				MaxScore: question.MaxScore,
				Results:  results,
				LoLinks:  loLinks,
			})
		}
		quizzes = append(quizzes, &model.Quiz{
			ID:        quiz.ID,
			Name:      quiz.Name,
			CreatedAt: quiz.CreatedAt,
			Questions: questions,
		})
	}
	return quizzes, nil
}

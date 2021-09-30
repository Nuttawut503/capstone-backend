package handler

import (
	"backend/database/question"
	"backend/database/questionlink"
	"backend/database/questionresult"
	"backend/database/quiz"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func getQuizHandlers() *fiber.App {
	app := fiber.New()

	app.Get("/quizzes", func(c *fiber.Ctx) error {
		type Question struct {
			QuestionID string `json:"questionID"`
			Title      string `json:"title"`
			MaxScore   int    `json:"maxScore"`
		}
		type Quiz struct {
			QuizID    string     `json:"quizID"`
			Title     string     `json:"title"`
			Questions []Question `json:"questions"`
		}
		quizzes := make([]Quiz, 0)
		simplifiedQuiz := quiz.GetQuizzesByCourseID(c.Query("courseID"))
		for _, q := range simplifiedQuiz {
			questions := make([]Question, 0)
			for _, v := range question.GetQuestionsByQuizID(q.QuizID) {
				questions = append(questions, Question{
					QuestionID: v.QuestionID,
					Title:      v.Title,
					MaxScore:   v.MaxScore,
				})
			}
			quizzes = append(quizzes, Quiz{
				QuizID:    q.QuizID,
				Title:     q.Title,
				Questions: questions,
			})
		}
		if len(quizzes) == 0 {
			quizzes = make([]Quiz, 0)
		}
		return c.JSON(quizzes)
	})

	app.Post("/quiz", func(c *fiber.Ctx) error {
		var newQuiz struct {
			CourseID string `json:"courseID"`
			Title    string `json:"title"`
		}
		if err := c.BodyParser(&newQuiz); err != nil {
			return err
		}
		quiz.AddRecord(newQuiz.CourseID, newQuiz.Title)
		return c.SendStatus(fiber.StatusCreated)
	})

	app.Post("/questions", func(c *fiber.Ctx) error {
		var newResult struct {
			QuizID    string `json:"quizID"`
			Questions string `json:"questions"`
		}
		if err := c.BodyParser(&newResult); err != nil {
			return err
		}
		type Question struct {
			Title        string `json:"title"`
			Maxscore     int    `json:"maxScore"`
			StudentID    int    `json:"studentID"`
			StudentScore int    `json:"studentScore"`
		}
		var newQuestions []Question
		if err := json.Unmarshal([]byte(newResult.Questions), &newQuestions); err != nil {
			return err
		}
		questionTitle := map[string]string{}
		for _, q := range newQuestions {
			questionID, added := questionTitle[q.Title]
			if !added {
				questionID := question.AddRecord(newResult.QuizID, q.Title, q.Maxscore)
				questionTitle[q.Title] = questionID
			}
			questionresult.AddRecord(questionID, q.StudentID, q.StudentScore)
		}
		return c.SendStatus(fiber.StatusCreated)
	})

	app.Get("/questionlinks", func(c *fiber.Ctx) error {
		// d
		// o
		c.Query("questionID")
		// t
		// h
		// i
		// s
		return c.JSON(quiz.GetQuizzesByCourseID(c.Query("courseID")))
	})

	app.Post("/questionlink", func(c *fiber.Ctx) error {
		var newQuestionLink struct {
			QuestionID string `json:"questionID"`
			LOID       string `json:"loID"`
			LOLevel    int    `json:"loLevel,int"`
		}
		if err := c.BodyParser(&newQuestionLink); err != nil {
			return err
		}
		questionlink.AddRecord(newQuestionLink.QuestionID, newQuestionLink.LOID, newQuestionLink.LOLevel)
		return c.SendStatus(fiber.StatusCreated)
	})

	return app
}

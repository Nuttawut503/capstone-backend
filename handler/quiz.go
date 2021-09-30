package handler

import (
	"backend/database/question"
	"backend/database/questionlink"
	"backend/database/questionresult"
	"backend/database/quiz"

	"github.com/gofiber/fiber/v2"
)

func getQuizHandlers() *fiber.App {
	app := fiber.New()

	app.Get("/quizzes", func(c *fiber.Ctx) error {
		// d
		// o

		// t
		// h
		// i
		// s
		return c.JSON(quiz.GetQuizzesByCourseID(c.Params("courseID", "")))
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
			Questions []struct {
				Title        string `json:"title"`
				Maxscore     int    `json:"maxscore"`
				StudentID    string `json:"studentID"`
				StudentScore int    `json:"studentScore,int"`
			} `json:"questions"`
		}
		if err := c.BodyParser(&newResult); err != nil {
			return err
		}
		questionTitle := map[string]string{}
		for _, q := range newResult.Questions {
			questionID, added := questionTitle[q.Title]
			if !added {
				question.AddRecord(newResult.QuizID, q.Title, q.Maxscore)
			}
			questionresult.AddRecord(questionID, q.StudentID, q.StudentScore)
		}
		return c.SendStatus(fiber.StatusCreated)
	})

	app.Get("/questionlinks", func(c *fiber.Ctx) error {
		// d
		// o
		c.Params("questionID", "")
		// t
		// h
		// i
		// s
		return c.JSON(quiz.GetQuizzesByCourseID(c.Params("courseID", "")))
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

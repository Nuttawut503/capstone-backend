package handler

import (
	"backend/database/course"
	"backend/database/lo"
	"backend/database/lolevel"
	"backend/database/plo"
	"backend/database/plolink"
	"backend/database/ploversion"
	"backend/database/program"
	"backend/database/question"
	"backend/database/questionlink"
	"backend/database/questionresult"
	"backend/database/quiz"
	"backend/database/student"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type SubLOResponseType struct {
	Level int    `json:"level,string"`
	Info  string `json:"info"`
}

type LOResponseType struct {
	LoID   string              `json:"loID"`
	Info   string              `json:"info"`
	Levels []SubLOResponseType `json:"levels"`
}

func GetHandler() *fiber.App {

	app := fiber.New()

	app.Get("/programs", func(c *fiber.Ctx) error {
		programs := program.GetRecords()
		if len(programs) == 0 {
			programs = make(program.Programs, 0)
		}
		return c.JSON(programs)
	})

	app.Post("/program", func(c *fiber.Ctx) error {
		var newProgram struct {
			Name string `json:"name"`
		}
		if err := c.BodyParser(&newProgram); err != nil {
			return err
		}
		program.AddRecord(newProgram.Name)
		return c.SendStatus(fiber.StatusCreated)
	})

	app.Get("/plos", func(c *fiber.Ctx) error {
		ploIDs := plo.GetPLOIDsByProgranID(c.Query("programID"))
		plos := ploversion.GetRecords(ploIDs)
		return c.JSON(plos)
	})

	app.Post("/plo", func(c *fiber.Ctx) error {
		var newPLO struct {
			ProgramID string `json:"programID"`
			Info      string `json:"info"`
		}
		if err := c.BodyParser(&newPLO); err != nil {
			return err
		}
		ploID := plo.AddRecord(newPLO.ProgramID)
		ploversion.AddRecord(ploID, newPLO.Info, 1)
		return c.SendStatus(fiber.StatusCreated)
	})

	app.Get("/courses", func(c *fiber.Ctx) error {
		courses := course.GetRecords(c.Query("programID"))
		if len(courses) == 0 {
			courses = make(course.Courses, 0)
		}
		return c.JSON(courses)
	})

	app.Post("/course", func(c *fiber.Ctx) error {
		var newCourse struct {
			ProgramID string `json:"programID"`
			Name      string `json:"name"`
			Semester  int    `json:"semester,string"`
			Year      int    `json:"year,string"`
		}
		if err := c.BodyParser(&newCourse); err != nil {
			fmt.Println(err)
			return err
		}
		course.AddRecord(newCourse.ProgramID, newCourse.Name, newCourse.Semester, newCourse.Year)
		return c.SendStatus(fiber.StatusCreated)
	})

	app.Get("/students", func(c *fiber.Ctx) error {
		students := student.GetRecords(c.Query("courseID"))
		if len(students) == 0 {
			students = make(student.Students, 0)
		}
		return c.JSON(students)
	})

	app.Post("/students", func(c *fiber.Ctx) error {
		var newStudents struct {
			CourseID string            `json:"courseID"`
			Students []student.Student `json:"students"`
		}
		if err := c.BodyParser(&newStudents); err != nil {
			return err
		}
		for i := range newStudents.Students {
			newStudents.Students[i].CourseID = newStudents.CourseID
		}
		student.AddRecords(newStudents.Students)
		return c.SendStatus(fiber.StatusCreated)
	})

	app.Get("/los", func(c *fiber.Ctx) error {
		var los []LOResponseType
		simplelos := lo.GetLOsByCourseID(c.Query("CourseID"))
		for _, simplelo := range simplelos {
			levels := make([]SubLOResponseType, 0)
			for _, slo := range lolevel.GetLevelInfoByLOID(simplelo[0]) {
				levels = append(levels, SubLOResponseType{
					Level: slo[0].(int),
					Info:  slo[1].(string),
				})
			}
			los = append(los, LOResponseType{
				LoID:   simplelo[0],
				Info:   simplelo[1],
				Levels: levels,
			})
		}
		return c.JSON(los)
	})

	app.Post("/lo", func(c *fiber.Ctx) error {
		var newLO struct {
			CourseID string `json:"courseID"`
			Info     string `json:"info"`
		}
		if err := c.BodyParser(&newLO); err != nil {
			return err
		}
		lo.AddRecord(newLO.CourseID, newLO.Info)
		return c.SendStatus(fiber.StatusCreated)
	})

	app.Post("/lolevel", func(c *fiber.Ctx) error {
		var newLOLevel struct {
			LOID  string `json:"loID"`
			Info  string `json:"info"`
			Level int    `json:"level,string"`
		}
		if err := c.BodyParser(&newLOLevel); err != nil {
			return err
		}
		lolevel.AddRecord(newLOLevel.LOID, newLOLevel.Info, newLOLevel.Level)
		return c.SendStatus(fiber.StatusCreated)
	})

	app.Post("/plolink", func(c *fiber.Ctx) error {
		var newPLOLink struct {
			PLOID string `json:"ploID"`
			LOID  string `json:"loID"`
		}
		if err := c.BodyParser(&newPLOLink); err != nil {
			return err
		}
		plolink.AddRecord(newPLOLink.PLOID, newPLOLink.LOID)
		return c.SendStatus(fiber.StatusCreated)
	})

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

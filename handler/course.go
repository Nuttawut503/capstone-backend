package handler

import (
	"backend/database/course"
	"backend/database/lo"
	"backend/database/lolevel"
	"backend/database/plolink"
	"backend/database/student"
	"encoding/json"

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

func getCourseHandlers() *fiber.App {
	app := fiber.New()

	app.Get("/courses", func(c *fiber.Ctx) error {
		// CourseID   string `json:"courseID"`
		// ProgramID  string `json:"programID"`
		// CourseName string `json:"courseName"`
		// Semester   int    `json:"semester"`
		// Year       int    `json:"year"`
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
			return err
		}
		course.AddRecord(newCourse.ProgramID, newCourse.Name, newCourse.Semester, newCourse.Year)
		return c.SendStatus(fiber.StatusCreated)
	})

	app.Get("/students", func(c *fiber.Ctx) error {
		// CourseID    string `json:"courseID"`
		// StudentID   string `json:"studentID"`
		// StudentName string `json:"studentName"`
		students := student.GetRecords(c.Query("courseID"))
		if len(students) == 0 {
			students = make(student.Students, 0)
		}
		return c.JSON(students)
	})

	app.Post("/students", func(c *fiber.Ctx) error {
		var newStudents struct {
			CourseID string `json:"courseID"`
			Students string `json:"students"`
		}
		if err := c.BodyParser(&newStudents); err != nil {
			return err
		}
		var students []student.Student
		if err := json.Unmarshal([]byte(newStudents.Students), &students); err != nil {
			return err
		}
		for i := range students {
			students[i].CourseID = newStudents.CourseID
		}
		student.AddRecords(students)
		return c.SendStatus(fiber.StatusCreated)
	})

	app.Get("/los", func(c *fiber.Ctx) error {
		// LoID   string              `json:"loID"`
		// Info   string              `json:"info"`
		// Levels []SubLOResponseType `json:"levels"`
		// Level int    `json:"level,string"`
		// Info  string `json:"info"`
		var los []LOResponseType
		simplelos := lo.GetLOsByCourseID(c.Query("courseID"))
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
		if len(los) == 0 {
			los = make([]LOResponseType, 0)
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

	return app
}

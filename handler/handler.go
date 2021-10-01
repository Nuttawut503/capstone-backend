package handler

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetHandlers() *fiber.App {

	app := fiber.New()

	db := Database{
		programs: map[string]Program{},
	}

	app.Get("/programs", func(c *fiber.Ctx) error {
		type ResponseFormat struct {
			ProgramID          string `json:"programID"`
			ProgramName        string `json:"programName"`
			ProgramDescription string `json:"programDescription"`
		}
		response := make([]ResponseFormat, 0)
		for k, v := range db.programs {
			response = append(response, ResponseFormat{
				ProgramID:          k,
				ProgramName:        v.programName,
				ProgramDescription: v.programDescription,
			})
		}
		return c.JSON(response)
	})

	app.Post("/program", func(c *fiber.Ctx) error {
		var request struct {
			ProgramName        string `json:"programName"`
			ProgramDescription string `json:"programDescription"`
		}
		if err := c.BodyParser(&request); err != nil {
			return err
		}
		db.createNewProgram(request.ProgramName, request.ProgramDescription)
		return c.SendStatus(fiber.StatusCreated)
	})

	app.Get("/courses", func(c *fiber.Ctx) error {
		programID := c.Query("programID")
		type ResponseFormat struct {
			CourseID   string `json:"courseID"`
			CourseName string `json:"courseName"`
			Semester   int    `json:"semester"`
			Year       int    `json:"year"`
		}
		response := make([]ResponseFormat, 0)
		for k, v := range db.programs[programID].courses {
			response = append(response, ResponseFormat{
				CourseID:   k,
				CourseName: v.courseName,
				Semester:   v.semester,
				Year:       v.year,
			})
		}
		return c.JSON(response)
	})

	app.Post("/course", func(c *fiber.Ctx) error {
		var request struct {
			ProgramID  string `json:"programID"`
			CourseName string `json:"courseName"`
			Semester   int    `json:"semester,string"`
			Year       int    `json:"year,string"`
		}
		if err := c.BodyParser(&request); err != nil {
			fmt.Println(err)
			return err
		}
		db.createNewCourse(request.ProgramID, request.CourseName, request.Semester, request.Year)
		return c.SendStatus(fiber.StatusCreated)
	})

	app.Get("/plos", func(c *fiber.Ctx) error {
		programID := c.Query("programID")
		type ResponseFormat struct {
			PLOID          string `json:"ploID"`
			PLOName        string `json:"ploName"`
			PLODescription string `json:"ploDescription"`
		}
		response := make([]ResponseFormat, 0)
		for k, v := range db.programs[programID].plos {
			response = append(response, ResponseFormat{
				PLOID:          k,
				PLOName:        v.ploName,
				PLODescription: v.ploDescription,
			})
		}
		return c.JSON(response)
	})

	app.Post("/plo", func(c *fiber.Ctx) error {
		var request struct {
			ProgramID      string `json:"programID"`
			PLOName        string `json:"ploName"`
			PLODescription string `json:"ploDescription"`
		}
		if err := c.BodyParser(&request); err != nil {
			return err
		}
		db.createNewPLO(request.ProgramID, request.PLOName, request.PLODescription)
		return c.SendStatus(fiber.StatusCreated)
	})

	app.Get("/students", func(c *fiber.Ctx) error {
		programID, courseID := c.Query("programID"), c.Query("courseID")
		type ResponseFormat struct {
			StudentID      string `json:"studentID"`
			StudentEmail   string `json:"studentEmail"`
			StudentName    string `json:"studentName"`
			StudentSurname string `json:"studentSurname"`
		}
		response := make([]ResponseFormat, 0)
		for k, v := range db.programs[programID].courses[courseID].students {
			response = append(response, ResponseFormat{
				StudentID:      k,
				StudentEmail:   v.studentEmail,
				StudentName:    v.studentName,
				StudentSurname: v.studentSurname,
			})
		}
		return c.JSON(response)
	})

	app.Post("/students", func(c *fiber.Ctx) error {
		var request struct {
			ProgramID string `json:"programID"`
			CourseID  string `json:"courseID"`
			Students  string `json:"students"`
		}
		if err := c.BodyParser(&request); err != nil {
			return err
		}
		var students []struct {
			StudentID      int    `json:"studentID"`
			StudentEmail   string `json:"studentEmail"`
			StudentName    string `json:"studentName"`
			StudentSurname string `json:"studentSurname"`
		}
		if err := json.Unmarshal([]byte(request.Students), &students); err != nil {
			fmt.Println(err)
			return err
		}
		for _, v := range students {
			db.addNewStudent(request.ProgramID, request.CourseID, strconv.Itoa(v.StudentID), v.StudentEmail, v.StudentName, v.StudentSurname)
		}
		return c.SendStatus(fiber.StatusCreated)
	})

	app.Get("/los", func(c *fiber.Ctx) error {
		programID, courseID := c.Query("programID"), c.Query("courseID")
		type (
			LevelFormat struct {
				Level            int    `json:"level"`
				LevelDescription string `json:"levelDescription"`
			}
			LinkedPLOFormat struct {
				PLOID   string `json:"ploID"`
				PLOName string `json:"ploName"`
			}
			ResponseFormat struct {
				LOID       string            `json:"loID"`
				LOTitle    string            `json:"loTitle"`
				Levels     []LevelFormat     `json:"levels"`
				LinkedPLOs []LinkedPLOFormat `json:"linkedPLOs"`
			}
		)
		response := make([]ResponseFormat, 0)
		for k, v := range db.programs[programID].courses[courseID].los {
			levels := make([]LevelFormat, 0)
			for _, v2 := range v.levels {
				levels = append(levels, LevelFormat{
					Level:            v2.level,
					LevelDescription: v2.levelDescription,
				})
			}
			linkedPLOs := make([]LinkedPLOFormat, 0)
			for k3 := range v.linkedploIDs {
				linkedPLOs = append(linkedPLOs, LinkedPLOFormat{
					PLOID:   k3,
					PLOName: db.programs[programID].plos[k3].ploName,
				})
			}
			response = append(response, ResponseFormat{
				LOID:       k,
				LOTitle:    v.loTitle,
				Levels:     levels,
				LinkedPLOs: linkedPLOs,
			})
		}
		return c.JSON(response)
	})

	app.Post("/lo", func(c *fiber.Ctx) error {
		var request struct {
			ProgramID   string `json:"programID"`
			CourseID    string `json:"courseID"`
			LOTitle     string `json:"loTitle"`
			InitLevel   int    `json:"initLevel,string"`
			Description string `json:"description"`
		}
		if err := c.BodyParser(&request); err != nil {
			return err
		}
		db.addNewLO(request.ProgramID, request.CourseID, request.LOTitle, request.InitLevel, request.Description)
		return c.SendStatus(fiber.StatusCreated)
	})

	app.Post("/lolevel", func(c *fiber.Ctx) error {
		var request struct {
			ProgramID   string `json:"programID"`
			CourseID    string `json:"courseID"`
			LOID        string `json:"loID"`
			Level       int    `json:"level,string"`
			Description string `json:"description"`
		}
		if err := c.BodyParser(&request); err != nil {
			return err
		}
		db.addNewLOLevel(request.ProgramID, request.CourseID, request.LOID, request.Level, request.Description)
		return c.SendStatus(fiber.StatusCreated)
	})

	app.Post("plolink", func(c *fiber.Ctx) error {
		var request struct {
			ProgramID string `json:"programID"`
			CourseID  string `json:"courseID"`
			PLOID     string `json:"PLOID"`
			LOID      string `json:"loID"`
		}
		if err := c.BodyParser(&request); err != nil {
			return err
		}
		db.addPLOLink(request.ProgramID, request.CourseID, request.PLOID, request.LOID)
		return c.SendStatus(fiber.StatusCreated)
	})

	return app
}

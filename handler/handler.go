package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetHandlers() *fiber.App {

	app := fiber.New()

	db := Database{
		programs: mockupProgram(),
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

	app.Get("/program-name", func(c *fiber.Ctx) error {
		programID := c.Query("programID")
		if _, ok := db.programs[programID]; !ok {
			return errors.New("wrong id")
		}
		response := struct {
			ProgramName string `json:"programName"`
		}{
			ProgramName: db.programs[programID].programName,
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

	app.Get("/course-name", func(c *fiber.Ctx) error {
		programID, courseID := c.Query("programID"), c.Query("courseID")
		if _, ok := db.programs[programID]; !ok {
			return errors.New("wrong id")
		}
		if _, ok := db.programs[programID].courses[courseID]; !ok {
			return errors.New("wrong id")
		}
		response := struct {
			CourseName string `json:"courseName"`
		}{
			CourseName: db.programs[programID].courses[courseID].courseName,
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

	app.Get("quizzes", func(c *fiber.Ctx) error {
		programID, courseID := c.Query("programID"), c.Query("courseID")
		type (
			LinkedLOFormat struct {
				LOID             string `json:"loID"`
				Level            int    `json:"level"`
				LevelDescription string `json:"levelDescription"`
			}
			QuestionFormat struct {
				QuestionID    string           `json:"questionID"`
				QuestionTitle string           `json:"questionTitle"`
				MaxScore      int              `json:"maxScore"`
				LinkedLOs     []LinkedLOFormat `json:"linkedLOs"`
			}
			ResponseFormat struct {
				QuizID    string           `json:"quizID"`
				QuizName  string           `json:"quizName"`
				Questions []QuestionFormat `json:"questions"`
			}
		)
		response := make([]ResponseFormat, 0)
		for quizID, quiz := range db.programs[programID].courses[courseID].quizzes {
			questionFormats := make([]QuestionFormat, 0)
			for questionID, question := range quiz.questions {
				linkedLOs := make([]LinkedLOFormat, 0)
				for loID, level := range question.linkedloIDs {
					linkedLOs = append(linkedLOs, LinkedLOFormat{
						LOID:             loID,
						Level:            level,
						LevelDescription: db.programs[programID].courses[courseID].los[loID].levels[level-1].levelDescription,
					})
				}
				questionFormats = append(questionFormats, QuestionFormat{
					QuestionID:    questionID,
					QuestionTitle: question.questionTitle,
					MaxScore:      question.maxScore,
					LinkedLOs:     linkedLOs,
				})
			}
			response = append(response, ResponseFormat{
				QuizID:    quizID,
				QuizName:  quiz.quizName,
				Questions: questionFormats,
			})
		}
		return c.JSON(response)
	})

	app.Post("quiz", func(c *fiber.Ctx) error {
		var request struct {
			ProgramID string `json:"programID"`
			CourseID  string `json:"courseID"`
			QuizName  string `json:"quizName"`
		}
		if err := c.BodyParser(&request); err != nil {
			return err
		}
		db.addQuiz(request.ProgramID, request.CourseID, request.QuizName)
		return c.SendStatus(fiber.StatusCreated)
	})

	app.Post("questions", func(c *fiber.Ctx) error {
		var request struct {
			ProgramID string `json:"programID"`
			CourseID  string `json:"courseID"`
			QuizID    string `json:"quizID"`
			Questions string `json:"questions"`
		}
		if err := c.BodyParser(&request); err != nil {
			return err
		}
		var newQuestions []QuestionExcel
		if err := json.Unmarshal([]byte(request.Questions), &newQuestions); err != nil {
			return err
		}
		db.addNewQuestion(request.ProgramID, request.CourseID, request.QuizID, newQuestions)
		return c.SendStatus(fiber.StatusCreated)
	})

	app.Post("questionlink", func(c *fiber.Ctx) error {
		var request struct {
			ProgramID  string `json:"programID"`
			CourseID   string `json:"courseID"`
			QuizID     string `json:"quizID"`
			QuestionID string `json:"questionID"`
			LOID       string `json:"loID"`
			Level      int    `json:"level,int"`
		}
		if err := c.BodyParser(&request); err != nil {
			fmt.Println(err)
			return err
		}
		db.addLOLink(request.ProgramID, request.CourseID, request.QuizID, request.QuestionID, request.LOID, request.Level)
		return c.SendStatus(fiber.StatusCreated)
	})

	app.Get("dashboard-flat", func(c *fiber.Ctx) error {
		programID, courseID := c.Query("programID"), c.Query("courseID")
		type (
			ResultFormat struct {
				StudentID    string `json:"studentID"`
				StudentScore int    `json:"studentScore"`
			}
			QuestionFormat struct {
				QuestionTitle string         `json:"questionTitle"`
				MaxScore      int            `json:"maxScore"`
				LinkedPLOs    []string       `json:"linkedPLOs"`
				LinkedLOs     []string       `json:"linkedLOs"`
				Results       []ResultFormat `json:"results"`
			}
			ResponseFormat struct {
				Student   map[string]string `json:"students"`
				PLOs      map[string]string `json:"plos"`
				LOs       map[string]string `json:"los"`
				Questions []QuestionFormat  `json:"questions"`
			}
		)
		response := ResponseFormat{
			Student:   map[string]string{},
			PLOs:      map[string]string{},
			LOs:       map[string]string{},
			Questions: []QuestionFormat{},
		}
		for _, quiz := range db.programs[programID].courses[courseID].quizzes {
			for _, question := range quiz.questions {
				plos := map[string]bool{}
				los := map[string]bool{}
				linkedLOs := []string{}
				linkedPLOs := []string{}
				results := []ResultFormat{}
				for loID, level := range question.linkedloIDs {
					response.LOs[loID] = db.programs[programID].courses[courseID].los[loID].loTitle

					for ploID := range db.programs[programID].courses[courseID].los[loID].linkedploIDs {
						plos[ploID] = true
						response.PLOs[ploID] = db.programs[programID].plos[ploID].ploName
					}

					if _, added := los[loID+","+strconv.Itoa(level)]; !added {
						los[loID+","+strconv.Itoa(level)] = true
						response.LOs[loID+","+strconv.Itoa(level)] = db.programs[programID].courses[courseID].los[loID].levels[level-1].levelDescription
						linkedLOs = append(linkedLOs, loID+","+strconv.Itoa(level))
					}
				}
				for ploID := range plos {
					linkedPLOs = append(linkedPLOs, ploID)
				}
				for _, student := range question.results {
					results = append(results, ResultFormat{
						StudentID:    student.studentID,
						StudentScore: student.studentScore,
					})
				}
				response.Questions = append(response.Questions, QuestionFormat{
					QuestionTitle: question.questionTitle,
					MaxScore:      question.maxScore,
					LinkedPLOs:    linkedPLOs,
					LinkedLOs:     linkedLOs,
					Results:       results,
				})
			}
		}
		for studentID, student := range db.programs[programID].courses[courseID].students {
			response.Student[studentID] = student.studentName
		}
		return c.JSON(response)
	})

	app.Get("dashboard-result", func(c *fiber.Ctx) error {
		programID, courseID := c.Query("programID"), c.Query("courseID")
		type (
			ResultFormat struct {
				StudentID    string `json:"studentID"`
				StudentName  string `json:"studentName"`
				StudentScore int    `json:"studentScore"`
			}
			ResponseFormat struct {
				QuizName string         `json:"quizName"`
				MaxScore int            `json:"maxScore"`
				Results  []ResultFormat `json:"results"`
			}
		)
		response := make([]ResponseFormat, 0)
		for _, quiz := range db.programs[programID].courses[courseID].quizzes {
			maxScore := 0
			studentScore := map[string]int{}
			for _, question := range quiz.questions {
				maxScore += question.maxScore
				for _, result := range question.results {
					if _, added := studentScore[result.studentID]; !added {
						studentScore[result.studentID] = result.studentScore
					} else {
						studentScore[result.studentID] += result.studentScore
					}
				}
			}
			result := make([]ResultFormat, 0)
			for studentID, score := range studentScore {
				studentName := "None"
				if student, ok := db.programs[programID].courses[courseID].students[studentID]; ok {
					studentName = student.studentName
				}
				result = append(result, ResultFormat{
					StudentID:    studentID,
					StudentName:  studentName,
					StudentScore: score,
				})
			}
			response = append(response, ResponseFormat{
				QuizName: quiz.quizName,
				MaxScore: maxScore,
				Results:  result,
			})
		}
		return c.JSON(response)
	})

	return app
}

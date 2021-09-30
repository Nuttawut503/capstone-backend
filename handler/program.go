package handler

import (
	"backend/database/plo"
	"backend/database/ploversion"
	"backend/database/program"

	"github.com/gofiber/fiber/v2"
)

func getProgramHandlers() *fiber.App {
	app := fiber.New()

	app.Get("/programs", func(c *fiber.Ctx) error {
		// ProgramID   string `json:"programID"`
		// UserID      string `json:"userID"`
		// ProgramName string `json:"programName"`
		// Description string `json:"description"`
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
		// PloID   string `json:"ploID"`
		// Info    string `json:"info"`
		// version int
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

	return app
}

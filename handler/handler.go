package handler

import (
	"github.com/gofiber/fiber/v2"
)

func GetHandlers() *fiber.App {

	app := fiber.New()

	app.Mount("/", getProgramHandlers())
	app.Mount("/", getCourseHandlers())
	app.Mount("/", getQuizHandlers())

	return app
}

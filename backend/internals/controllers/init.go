package controllers

import (
	"github.com/gofiber/fiber/v2"
)

// var CsrfMiddleware fiber.Handler

func InitializeCintrollers(app *fiber.App) {
	AuthController(app)
	EventsController(app)
	app.Get("test", func(c *fiber.Ctx) error {
		// configs.AppDB.Ping()
		return c.JSON(fiber.Map{
			"message": "app is working",
		})
	})
}

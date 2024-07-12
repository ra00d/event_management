package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ra00d/event_management/internals/configs"
	"github.com/ra00d/event_management/internals/handlers"
)

func AuthController(app *fiber.App) {
	authGroup := app.Group("auth")
	authGroup.Post("sign-up", handlers.SignUp)
	authGroup.Post("log-in", handlers.Login)
	authGroup.Options("log-in", handlers.Login)
	authGroup.Get("log-out", func(c *fiber.Ctx) error {
		sess, err := configs.GetSessionStore(c)
		if err != nil {
			return c.Status(500).JSON(err)
		}
		sess.Destroy()
		return c.SendStatus(200)
	})
}

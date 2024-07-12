package middelwares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

// const HeaderName = "X-Csrf-Token"

func MiddelwareInit(app *fiber.App) {
	app.Use(
		cors.New(
			cors.Config{
				AllowOrigins:     "http://localhost:4200, http://localhost:5173,http://localhost:8080",
				AllowCredentials: true,
			},
		),
	)
	// app.Use(helmet.New())
	app.Use(recover.New())
	// app.Get("/csrf", configs.CsrfMiddleware(), func(c *fiber.Ctx) error {
	// 	csrfToken, ok := c.Locals("csrf").(string)
	// 	if !ok {
	// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 			"error":   fiber.ErrInternalServerError,
	// 			"message": "can not set token",
	// 		})
	// 	}
	//
	// 	return c.JSON(fiber.Map{
	// 		"csrf": csrfToken,
	// 	})
	// })
	// app.Use(configs.CsrfMiddleware())
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format: "${locals:requestid} ${status} - ${method} ${path}\u200b\n",
	}))
	SetDocumntationConfig(app)
}

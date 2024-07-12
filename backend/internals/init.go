package internals

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ra00d/event_management/internals/configs"
	"github.com/ra00d/event_management/internals/controllers"
	"github.com/ra00d/event_management/internals/middelwares"
)

func init() {
	configs.DataBasaInit()
	configs.SessionsInit()
}

func Init() {
	app := fiber.New(configs.FiberConfig())
	middelwares.MiddelwareInit(app)
	// app.Static("/", "./public")
	// app.Static("/assets", "./public")
	// Setup static files
	controllers.InitializeCintrollers(app)
	app.Static("/public", "./public")
	app.Static("/uploads", "./storage/uploads")
	app.Listen(":8080")
}

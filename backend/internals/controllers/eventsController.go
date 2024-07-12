package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ra00d/event_management/internals/handlers"
)

func EventsController(app *fiber.App) {
	handler := handlers.EventsHandler{}
	group := app.Group("events")
	group.Get("", handler.GetEvents)
	group.Get(":id", handler.GetOne)
	group.Post("", handler.AddEvent)
	group.Delete(":id", handler.Delete)
	group.Put(":id", handler.Update)
}

package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ra00d/event_management/internals/models"
	"github.com/ra00d/event_management/internals/queries"
)

type handler interface {
	GetAll(ctx *fiber.Ctx) error
	GetOne(ctx *fiber.Ctx) error
	AddEvent(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
}
type EventsHandler struct{}

func (h EventsHandler) GetEvents(ctx *fiber.Ctx) error {
	return ctx.JSON(queries.GetAllEvents())
}

func (h EventsHandler) AddEvent(c *fiber.Ctx) error {
	event := models.AddEventBody{}
	err := c.BodyParser(&event)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   err,
			"message": "fialed to parse the body",
		})
	}
	if form, err := c.MultipartForm(); err == nil {
		files := form.File["images[]"]
		queries.AddEvent(event, files, c.SaveFile)
	}
	// queries
	return c.SendStatus(fiber.StatusCreated)
}

func (h EventsHandler) GetOne(c *fiber.Ctx) error {
	res := queries.GetEvent(c.Params("id", "0"))
	return c.JSON(res)
}

func (h EventsHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id", "0")
	res, err := queries.DeleteEvent(id)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.JSON(res)
}

func (h EventsHandler) Update(c *fiber.Ctx) error {
	event := models.UpdateEventBody{}

	err := c.BodyParser(&event)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   err,
			"message": "fialed to parse the body",
		})
	}

	err = event.Validate()
	if err != nil {
		return c.Status(422).JSON(err)
	}

	id := c.Params("id", "0")

	if form, err := c.MultipartForm(); err == nil {
		files := form.File["images[]"]
		res, _ := queries.UpdateEvent(id, event, files, c.SaveFile)
		return c.Status(fiber.StatusOK).JSON(res)
	}
	return c.Status(500).JSON(fiber.Map{
		"error":   err,
		"message": "fialed to parse the body",
	})
}

package configs

import (
	// "errors"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	exceptions "github.com/ra00d/event_management/internals/constants/errors"
)

func FiberConfig() fiber.Config {
	return fiber.Config{
		AppName:      "Event Management Api",
		ErrorHandler: defaultErrorHandler,
		// EnablePrintRoutes: true,
		// Views:        AttachTemplateEngine(),
	}
}

func defaultErrorHandler(c *fiber.Ctx, err error) error {
	fmt.Println(err.Error())
	code := fiber.StatusInternalServerError

	var e *fiber.Error
	var notFoundErr *exceptions.NotFoundError
	if errors.As(err, &e) {
		fmt.Println(err.Error())
		code = e.Code
	}
	if errors.As(err, &notFoundErr) {
		fmt.Println(err.Error())
		code = notFoundErr.Code
	}
	c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

	return c.Status(code).JSON(err)
}

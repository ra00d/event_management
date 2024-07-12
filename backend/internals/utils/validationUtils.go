package utils

import (
	"github.com/gofiber/fiber/v2"
)

type Validator interface {
	Validate() error
}

func ValidateAndParse(payload Validator, c *fiber.Ctx) (int, fiber.Map) {
	var err error
	// fmt.Println(payload)
	err = c.BodyParser(&payload)
	if err != nil {
		return 500, fiber.Map{
			"message": "sorry something went wrong in our side",
			"error":   err.Error(),
		}
	}
	// fmt.Println(payload)
	err = payload.Validate()
	if err != nil {
		return 400, fiber.Map{
			"message":    "some fields are not valid",
			"errors":     err,
			"statusCode": 400,
		}
	}
	return 0, nil
}

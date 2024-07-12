package middelwares

import (
	"encoding/json"
	"slices"

	"github.com/gofiber/fiber/v2"
	"github.com/ra00d/event_management/internals/configs"
)

type AuthOptions struct {
	Role        int
	Permissions []int
}

func AuthMiddleware(opt AuthOptions) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := configs.GetSessionStore(c)
		if err != nil {
			return c.Status(500).JSON(err)
		}

		// check if the user is loged in if not retutn 403 unauthenticated
		userId := sess.Get("user_id")
		if userId == nil || userId == "" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"statusCode": 403,
				"message":    fiber.ErrForbidden.Message,
				"err":        fiber.ErrForbidden,
			})
		}
		// get the user role from session
		// if not exist return 401 unauthorized
		role := sess.Get("role")
		if role == "" || role != opt.Role {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"statusCode": 401,
				"message":    fiber.ErrUnauthorized.Message,
				"err":        fiber.ErrUnauthorized,
			})
		}
		// check the user permisions permissions
		// return 401 unauthriazed if does not macth
		permissions := sess.Get("permissions").([]byte)
		json.Unmarshal(permissions, &permissions)
		for _, v := range opt.Permissions {
			if !slices.Contains(permissions, byte(v)) {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"statusCode": 401,
					"message":    fiber.ErrUnauthorized.Message,
					"err":        fiber.ErrUnauthorized,
				})
			}
		}

		return c.Next()
	}
}

// func SetAuthData(role string, permissons []string) {
// }

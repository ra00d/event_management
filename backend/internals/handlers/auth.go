package handlers

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"

	// "github.com/jmoiron/sqlx"
	"github.com/ra00d/event_management/internals/configs"
	"github.com/ra00d/event_management/internals/constants/auth"
	"github.com/ra00d/event_management/internals/models"
	"github.com/ra00d/event_management/internals/queries"
	"github.com/ra00d/event_management/internals/utils"
)

func SignUp(c *fiber.Ctx) error {
	payload := models.SignUpModel{}
	var err error
	statusCode, errMap := utils.ValidateAndParse(&payload, c)
	if errMap != nil {
		return c.Status(statusCode).JSON(errMap)
	}

	payload.Password, err = utils.HashPassword(payload.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":    "internal server error",
			"statusCode": 500,
			"error":      err.Error(),
		})
	}
	err = queries.MustCreateUser(auth.USER, payload.Username, payload.Password, payload.Email)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.SendStatus(fiber.StatusCreated)
}

func Login(c *fiber.Ctx) error {
	// fmt.Println("in login")
	// token, ok := c.Locals("csrf").(string)
	// if !ok {
	// 	return c.Status(403).JSON(fiber.Map{
	// 		"csrf": token,
	// 	})
	// }
	user := models.User{}
	loginModel := models.LoginModel{}
	var err error
	err = c.BodyParser(&loginModel)
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// pass := loginModel.Password
	// user.IsActive = true
	err = configs.AppDB.Get(&user,
		`SELECT user_id as id,role_id as role,password_hash as password from users
	         	where email=? AND is_active=?`,
		loginModel.Email,
		true,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.Status(403).JSON(fiber.Map{
				"message": "wrong email or password",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if !utils.CompareHash(loginModel.Password, user.Password) {
		return c.Status(403).JSON(fiber.Map{
			"message": "wrong email or password",
		})
	}

	sess, err := configs.GetSessionStore(c)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	sess.Set("user_id", user.ID)

	res, err := configs.AppDB.Queryx(
		`SELECT json_arrayagg(permission_id) as permissions from user_permissions
		WHERE user_id = ? 
		GROUP BY user_id `,
		user.ID,
	)
	if err != nil {
		fmt.Println(err, "permissions")
		return c.Status(500).JSON(err)
	}
	result := map[string]interface{}{}
	for res.Next() {
		err := res.MapScan(result)
		if err != nil {
			fmt.Println(err, "map")
			return c.Status(500).JSON(res.Err())
		}

	}
	sess.Set("user_id", user.ID)
	sess.Set("role", user.Role)
	sess.Set("permissions", result["permissions"].([]byte))
	err = sess.Save()
	if err != nil {
		fmt.Println(err)
		fmt.Println("session")
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "loged-in",
		// "user":    result,
	})
}

package handler

import (
	"context"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/subhasbodaki/project_mgmt/db"
	"github.com/subhasbodaki/project_mgmt/postgres"
)

//get user details from DB
func GetUserByEmail(e string) (*postgres.GetEmailRow, error) {
	user, err := db.DB.GetEmail(context.Background(), e)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func Login(c *fiber.Ctx) error {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var user request
	err := c.BodyParser(&user) //It recieves the login details from Body and parse it
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse json",
		})
	}

	userData, err := GetUserByEmail(user.Email)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"messagae": "User Not Found..!",
		})
	}

	//match email & password
	if user.Email != userData.Email || user.Password != userData.Password {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	//create token
	token := jwt.New(jwt.SigningMethodHS256)

	//set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = "sbodaki@gmail.com"
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7) //Week

	//generate encoded token
	s, err := token.SignedString([]byte("secret"))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	//send encoded token as response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": s,
	})
}

func AuthRequired() func(c *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
		SigningKey: []byte("secret"),
	})
}

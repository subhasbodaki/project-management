package handler

import "github.com/gofiber/fiber/v2"

func Hello(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Hii, This is Project Management Server",
	})
}

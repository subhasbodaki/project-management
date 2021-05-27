package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/subhasbodaki/project_mgmt/src/routes"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "Server is up & running on port 3000",
		})
	})

	api := app.Group("/api")

	api.Get("", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "Server is up & running on port 3000",
		})
	})

	//Connet project_mgmt routes
	routes.ProjectRoutes(api.Group("/projects"))
}

func main() {
	app := fiber.New()

	setupRoutes(app)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}

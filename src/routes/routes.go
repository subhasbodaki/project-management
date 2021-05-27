package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/subhasbodaki/project_mgmt/src/services"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "Server is up & running on port 3000",
		})
	})

	api := app.Group("/api")
	//route:  /api/
	api.Get("", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "Server is up & running on port 3000",
		})
	})
	//route : /api/projects
	projects := api.Group("/projects")

	//route : /api/projects
	projects.Post("", services.CreateProject)
	projects.Get("", services.GetProjects)
	projects.Get("/:id", services.GetProjectsById)
	projects.Patch("/:id", services.UpdateProjectById)
	projects.Delete("/:id", services.DeleteProjectById)

}

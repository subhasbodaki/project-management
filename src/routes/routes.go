package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/subhasbodaki/project_mgmt/src/handler"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", handler.Hello)

	api := app.Group("/api")

	app.Post("/login", handler.Login)
	//route:  /api/
	api.Get("", handler.Hello)

	//route: /api/user
	api.Post("/user", handler.ResgisterUser)

	//route : /api/projects
	projects := api.Group("/projects")

	//route : /api/projects
	projects.Post("", handler.CreateProject)
	projects.Get("", handler.GetProjects)
	projects.Get("/:id", handler.GetProjectsById)
	projects.Patch("/:id", handler.UpdateProjectById)
	projects.Delete("/:id", handler.DeleteProjectById)

}

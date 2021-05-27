package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/subhasbodaki/project_mgmt/src/services"
)

func ProjectRoutes(route fiber.Router) {
	route.Post("", services.CreateProject)
	route.Get("", services.GetProjects)
	route.Get("/:id", services.GetProjectsById)
	route.Patch("/:id", services.UpdateProjectById)
	route.Delete("/:id", services.DeleteProjectById)
}

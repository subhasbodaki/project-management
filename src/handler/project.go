package handler

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/subhasbodaki/project_mgmt/db"
	"github.com/subhasbodaki/project_mgmt/postgres"
)

type Project struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	StartDate   string `json:"startdate"`
	EndDate     string `json:"enddate"`
	Active      bool   `json:"active"`
}

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ResgisterUser(c *fiber.Ctx) error {
	var u User

	if err := c.BodyParser(&u); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Unable to parse..!",
		})
	}

	user, err := db.DB.CreateUser(context.Background(), postgres.CreateUserParams{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Unable to register user..!",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}

func CreateProject(c *fiber.Ctx) error {
	p := new(Project)

	if err := c.BodyParser(p); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Unable to parse",
		})
	}

	project, err := db.DB.CreateProject(context.Background(), postgres.CreateProjectParams{
		Name:        p.Name,
		Description: p.Description,
		StartDate:   p.StartDate,
		EndDate:     p.EndDate,
		Active:      p.Active,
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Unable to add project",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(project)
}

func GetProjects(c *fiber.Ctx) error {
	project, err := db.DB.GetProjects(context.Background())
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Unable to process request",
		})
	}

	return c.Status(fiber.StatusOK).JSON(project)
}

func GetProjectsById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Fatal(err)
	}

	project, err := db.DB.GetProjectById(context.Background(), int32(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Project Not Found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(project)
}

func UpdateProjectById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Fatal(err)
	}
	p := new(Project)

	if err := c.BodyParser(p); err != nil {
		return err
	}

	project, err := db.DB.UpdateProjectById(context.Background(), postgres.UpdateProjectByIdParams{
		ID:     int32(id),
		Active: p.Active,
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Unable to update the project",
		})
	}

	return c.Status(fiber.StatusOK).JSON(project)
}

func DeleteProjectById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Fatal(err)
	}

	err = db.DB.DeleteProjectById(context.Background(), int32(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Project Not Found..!",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "successfully Deleted",
	})
}

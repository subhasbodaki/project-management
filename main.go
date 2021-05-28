package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/lib/pq"
	"github.com/subhasbodaki/project_mgmt/db"
	"github.com/subhasbodaki/project_mgmt/src/handler"
	"github.com/subhasbodaki/project_mgmt/src/routes"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Use("/api/*", handler.AuthRequired())

	db.DBConn()

	routes.SetupRoutes(app)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}

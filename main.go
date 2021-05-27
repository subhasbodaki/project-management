package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/subhasbodaki/project_mgmt/db"
	"github.com/subhasbodaki/project_mgmt/src/routes"
)

func main() {
	app := fiber.New()

	db.DBConn()

	routes.SetupRoutes(app)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}

package db

import (
	"database/sql"
	"log"

	"github.com/subhasbodaki/project_mgmt/postgres"
)

var DB *postgres.Queries

func DBConn() {
	conn, err := sql.Open("postgres", "postgresql://postgres:password@localhost:5432/project_management")

	if err != nil {
		log.Fatal(err)
	}

	DB = postgres.New(conn)
}

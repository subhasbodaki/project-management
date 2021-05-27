package sqldb

import (
	"database/sql"
	"log"

	"github.com/subhasbodaki/project_mgmt/postgres"
)

func dbconn() *postgres.Queries {
	conn, err := sql.Open("postgres", "postgresql://postgres:password@localhost:5432/project_management")

	if err != nil {
		log.Fatal(err)
	}

	db := postgres.New(conn)

	return db
}

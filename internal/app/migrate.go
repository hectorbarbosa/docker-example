package app

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log"
)

func MakeMigration(steps int, db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations/",
		"users_db", driver)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Steps(steps); err != nil {
		log.Println("steps: ", err)
	}
	// if err = m.Up(); err != nil {
	// 	log.Println("up: ", err)
	// }
}

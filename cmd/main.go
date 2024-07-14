package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"docker-example/internal/app"
)

func main() {
	//connect to database
	// db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	db, err := createConnection()
	if err != nil {
		log.Fatal("Failed to connect to database\n", err)
	}

	// migration up, create table users
	app.MakeMigration(1, db)
	app.PopulateDB(db)

	// id := 2
	// u := app.GetUser(id, db)
	// fmt.Println(u.Name, u.Email)

	count := app.GetCount(db)
	fmt.Println("number or rows: ", count)

	// migration up, create table users
	app.MakeMigration(-1, db)
}

func createConnection() (*sql.DB, error) {
	// connStr := fmt.Sprintf(
	// 	"postgres://%s:%s@db/%s?sslmode=disable",
	// 	// os.Getenv("DB_USER"),
	// 	// os.Getenv("DB_PASSWORD"),
	// 	// os.Getenv("DB_NAME"),
	// 	"postgres",
	// 	"postgres",
	// 	"users_db",
	// )
	// fmt.Println("connection string: ", connStr)
	// db, err := sql.Open("postgres", connStr)
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	return db, nil
}

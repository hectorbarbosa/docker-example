package app

import (
	"database/sql"
	"log"
)

func PopulateDB(db *sql.DB) {
	createUser(&User{Name: "User1", Email: "myemail@mail.com"}, db)
	createUser(&User{Name: "User2", Email: "myemail2@mail.com"}, db)
}

// create user
func createUser(u *User, db *sql.DB) *User {
	err := db.QueryRow(
		"INSERT INTO public.users (name, email) VALUES ($1, $2) RETURNING id;",
		u.Name, u.Email).Scan(&u.Id)
	if err != nil {
		log.Fatal(err)
	}

	return u
}

// get user by id
func GetUser(id int, db *sql.DB) User {
	var u User
	err := db.QueryRow("SELECT * FROM public.users WHERE id = $1;", id).Scan(&u.Id, &u.Name, &u.Email)
	if err != nil {
		log.Fatal(err)
	}
	return u
}

// get all records count
func GetCount(db *sql.DB) int {
	var count int
	err := db.QueryRow("SELECT count(*) as c FROM public.users;").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	return count
}

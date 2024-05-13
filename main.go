package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Student struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	Age      int    `db:"age"`
	Location string `db:"location"`
}

func main() {
	//connect to a PostgreSQL database
	// Replace the connection details (user, dbname, password, host) with your own
	db, err := sqlx.Connect("postgres", "user=postgres dbname=mani sslmode=disable password=postgres host=localhost")

	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	// Test the connection to the database
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully Connected")
	}
}

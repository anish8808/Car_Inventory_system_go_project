package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	dsn := "user=manager database=postgres sslmode=disable"

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		fmt.Println("Error opening database", err)
		panic(err)
	}

	if err := db.Ping(); err != nil {
		fmt.Println("Error opening database", err)
		panic(err)
	}

	fmt.Println("Successfully connected to the database")
	DB = db
}

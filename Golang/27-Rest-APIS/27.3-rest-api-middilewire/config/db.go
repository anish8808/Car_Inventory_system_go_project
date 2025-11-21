package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	dsn := "user=manager database=postgres sslmode=disable "

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println("Error openingn databse", err)
		panic(err)
	}

	if err := db.Ping(); err != nil {
		fmt.Println("Error connecting database ", err)
	}

	fmt.Println("Successfully connected to the databae ")

	DB = db
}

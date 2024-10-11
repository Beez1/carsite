package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL driver
)

type User struct {
	Name  string `db:"username"`
	Email string `db:"email"`
}

func main() {
	// Connection string
	dsn := "user=beez password=carsitetest1234 dbname=carsite host=35.246.1.155 port=5432 sslmode=disable"

	// Connect to the PostgreSQL database
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalln("Error connecting to the database:", err)
	}
	defer db.Close()

	// Test the connection
	if err := db.Ping(); err != nil {
		log.Fatal("Error pinging the database:", err)
	} else {
		log.Println("Successfully connected to the database")
	}
}

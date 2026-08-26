package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // blank import - registers the driver, we don't call it directly
)

func main() {
	//load the .env file
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("error loading .env file", err)
	}

	//read the pswd from the env variable
	password := os.Getenv("DB_PASSWORD")

	connStr := fmt.Sprintf(
		"host=localhost port=5432 user=postgres password=%s dbname=coursedb sslmode=disable",
		password,
	)

	//sql.Open doesnt actuakly connect yet, just prepares to connect
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("error opening db", err)
	}

	defer db.Close()

	//ping actually tests the connection
	err = db.Ping()
	if err != nil {
		log.Fatal("error connecting to db", err)
	}

	fmt.Println("succesfully connected to postgres !")
}

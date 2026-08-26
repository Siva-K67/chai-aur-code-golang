package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // blank import - registers the driver, we don't call it directly
)

// Course represents one row in our courses table
type Course struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

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

	//create the courses tabkle if it dosent exist already
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS courses(
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	price INT NOT NULL
	)`

	// db.Exec(...) — used for SQL statements that don't return rows
	// (CREATE, INSERT, UPDATE, DELETE)
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal("error creating table: ", err)
	}
	fmt.Println("courses table ready")

	//insert a new course
	// notice ID is left unset. because this PK is serial
	newCourse := Course{Name: "Boxing Guide by Ippo Makunochi", Price: 499}

	// $1 and $2 are placeholders
	insertQuery := `INSERT INTO courses (name,price) VALUES ($1,$2)`

	_, err = db.Exec(insertQuery, newCourse.Name, newCourse.Price)
	if err != nil {
		log.Fatal("error inserting course", err)
	}
	fmt.Println("course inserted succesfully")

}

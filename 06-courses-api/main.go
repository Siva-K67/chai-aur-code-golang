package main

import (
	"database/sql"
	"errors"
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

// function to retrive all the courses
func getAllCourses(db *sql.DB) ([]Course, error) {

	//rows doesn't hold all the data at once, sitting there ready to index into.
	// It's a live connection pointing at one row at a time, streaming from the
	// database as you ask for more

	// SELECT id, name, price (not SELECT *) - Scan matches columns by POSITION, not name.
	// If we used SELECT * and the table's column order ever changed (e.g. a new
	// column added in the middle), Scan would silently misalign or error out,
	// since it has no idea about column names - only "1st result -> 1st arg", etc
	rows, err := db.Query("SELECT id, name, price FROM courses")
	if err != nil {
		return nil, err
	}

	//close the result set when done
	defer rows.Close()

	var courses []Course

	// rows.Next() advances to the next row - returns false when there are none left
	for rows.Next() {
		var c Course
		// Scan copies the current row's columns into c's fields, in order given in Scan's args
		err := rows.Scan(&c.ID, &c.Name, &c.Price)
		if err != nil {
			return nil, err
		}
		courses = append(courses, c)
	}

	return courses, nil
}

// function to fetch a single course
func getCourseByID(db *sql.DB, id int) (Course, error) {
	var c Course

	// QueryRow expects exactly one row back - simpler than Query for a single record
	row := db.QueryRow("SELECT id, name, price FROM courses WHERE id=$1", id)

	err := row.Scan(&c.ID, &c.Name, &c.Price)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Course{}, fmt.Errorf("course with id %d not found: %w", id, err)
		}

		return Course{}, err
	}

	return c, nil
}

// function UPDATEs an existing course's name and price by ID
func updateCourse(db *sql.DB, id int, name string, price int) error {
	updateQuery := `UPDATE courses SET name = $1, price = $2 WHERE id = $3`

	result, err := db.Exec(updateQuery, name, price, id)
	if err != nil {
		return err
	}

	// RowsAffected tells us how many rows the UPDATE actually changed
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	// if 0 rows changed, that ID didn't exist
	if rowsAffected == 0 {
		return fmt.Errorf("course with id %d not found", id)
	}

	return nil
}

// Function to DELETE a course by ID - same RowsAffected pattern as update
func deleteCourse(db *sql.DB, id int) error {
	deleteQuery := `DELETE FROM courses WHERE id = $1`

	result, err := db.Exec(deleteQuery, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("course with id %d not found", id)
	}

	return nil
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

	// //insert a new course
	// // notice ID is left unset. because this PK is serial
	// newCourse := Course{Name: "Death NOte by Tsugumi Ohba", Price: 300}

	// // $1 and $2 are placeholders
	// insertQuery := `INSERT INTO courses (name,price) VALUES ($1,$2)`

	// _, err = db.Exec(insertQuery, newCourse.Name, newCourse.Price)
	// if err != nil {
	// 	log.Fatal("error inserting course", err)
	// }
	// fmt.Println("course inserted succesfully")

	//query all of the courses
	// courses, err := getAllCourses(db)
	// if err != nil {
	// 	log.Fatal("error fetching the courses:", err)
	// }

	// for _, c := range courses {
	// 	fmt.Println(c.ID, c.Name, c.Price)
	// }

	// // try fetching a course that should exist
	// course, err := getCourseByID(db, 5)
	// if err != nil {
	// 	fmt.Println("error:", err)
	// } else {
	// 	fmt.Println("found:", course.ID, course.Name, course.Price)
	// }

	// // try fetching one that shouldn't exist
	// course2, err2 := getCourseByID(db, 999)
	// if err2 != nil {
	// 	fmt.Println("error:", err2)
	// } else {
	// 	fmt.Println("found:", course2.ID, course2.Name, course2.Price)
	// }

	// // update course with id 2
	// err = updateCourse(db, 2, "One Piece by Takeshi Obada", 599)
	// if err != nil {
	// 	fmt.Println("error updating:", err)
	// } else {
	// 	fmt.Println("course updated successfully")
	// }

	// // try updating one that doesn't exist
	// err = updateCourse(db, 999, "Ghost Course", 0)
	// if err != nil {
	// 	fmt.Println("error updating:", err)
	// } else {
	// 	fmt.Println("course updated successfully")
	// }

	// delete a course that exists
	err = deleteCourse(db, 3)
	if err != nil {
		fmt.Println("error deleting:", err)
	} else {
		fmt.Println("course deleted successfully")
	}

	// try deleting one that doesn't exist
	err = deleteCourse(db, 777)
	if err != nil {
		fmt.Println("error deleting:", err)
	} else {
		fmt.Println("course deleted successfully")
	}

}

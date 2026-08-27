package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// function to retrive all the courses
func getAllCourses(db *sql.DB) ([]Course, error) {

	//rows doesn't hold all the data at once.
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

// inserts the course into the db
func insertCourse(name string, price int) error {
	_, err := db.Exec("INSERT INTO courses (name,price) VALUES ($1,$2)", name, price)
	return err
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

// courseHandler handler /courses api.
func coursesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		courses, err := getAllCourses(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(courses)

	case http.MethodPost:
		var newCourse Course
		if err := json.NewDecoder(r.Body).Decode(&newCourse); err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}
		if err := insertCourse(newCourse.Name, newCourse.Price); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "course created"})

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

// courseByIDHandler handles /courses/{id} - GET (one), PUT (update), DELETE
func courseByIDHandler(w http.ResponseWriter, r *http.Request) {
	// extract the id from the URL path, e.g. "/courses/3" -> "3"
	idStr := strings.TrimPrefix(r.URL.Path, "/courses/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		course, err := getCourseByID(db, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(course)

	case http.MethodPut:
		var updated Course
		if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}
		if err := updateCourse(db, id, updated.Name, updated.Price); err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "course updated"})

	case http.MethodDelete:
		if err := deleteCourse(db, id); err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "course deleted"})

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

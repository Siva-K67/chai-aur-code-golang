package main

// Course represents one row in our courses table
type Course struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

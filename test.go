package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStudents(t *testing.T) {
	// Setup
	app := setupApp()
	defer app.DB().Close()

	// Insert test data into the database
	insertTestData(app.DB())

	// Create a GET request to /students
	req, err := http.NewRequest("GET", "/students", nil)
	assert.NoError(t, err)

	// Execute the request
	resp := httptest.NewRecorder()
	app.ServeHTTP(resp, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, resp.Code)

	// Decode the response body
	var students []Students
	err = json.NewDecoder(resp.Body).Decode(&students)
	assert.NoError(t, err)

	// Check the response data
	assert.Len(t, students, 2)
	assert.Equal(t, "John", students[0].First_name)
	assert.Equal(t, "Doe", students[0].Last_name)
	assert.Equal(t, 25, students[0].Age)
	// Add more assertions based on your data
}

func TestCreateStudent(t *testing.T) {
	// Setup
	app := setupApp()
	defer app.DB().Close()

	// Create a POST request to add a new student
	req, err := http.NewRequest("POST", "/students/John/Doe/25/1/Male/Address/1234567890", nil)
	assert.NoError(t, err)

	// Execute the request
	resp := httptest.NewRecorder()
	app.ServeHTTP(resp, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, resp.Code)

	// Query the database to check if the new student is added
	var count int
	err = app.DB().QueryRow("SELECT COUNT(*) FROM Student_details WHERE first_name='John' AND last_name='Doe'").Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 1, count)
}

func TestUpdateStudent(t *testing.T) {
	// Setup
	app := setupApp()
	defer app.DB().Close()

	// Insert test data into the database
	insertTestData(app.DB())

	// Create a PUT request to update a student
	req, err := http.NewRequest("PUT", "/students/1/Updated/Name/30/2/Female/UpdatedAddress/9876543210", nil)
	assert.NoError(t, err)

	// Execute the request
	resp := httptest.NewRecorder()
	app.ServeHTTP(resp, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, resp.Code)

	// Query the database to check if the student is updated
	var updatedStudent Students
	err = app.DB().QueryRow("SELECT * FROM Student_details WHERE id=1").Scan(
		&updatedStudent.ID,
		&updatedStudent.First_name,
		&updatedStudent.Last_name,
		&updatedStudent.Age,
		&updatedStudent.Class,
		&updatedStudent.Gender,
		&updatedStudent.Address,
		&updatedStudent.Phone_number,
	)
	assert.NoError(t, err)

	// Check the updated data
	assert.Equal(t, "Updated", updatedStudent.First_name)
	assert.Equal(t, "Name", updatedStudent.Last_name)
	assert.Equal(t, 30, updatedStudent.Age)
	// Add more assertions based on your data
}

func TestDeleteStudent(t *testing.T) {
	// Setup
	app := setupApp()
	defer app.DB().Close()

	// Insert test data into the database
	insertTestData(app.DB())

	// Create a DELETE request to delete a student
	req, err := http.NewRequest("DELETE", "/students/1", nil)
	assert.NoError(t, err)

	// Execute the request
	resp := httptest.NewRecorder()
	app.ServeHTTP(resp, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, resp.Code)

	// Query the database to check if the student is deleted
	var count int
	err = app.DB().QueryRow("SELECT COUNT(*) FROM Student_details WHERE id=1").Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 0, count)
}

// Utility function to set up the Gofr app
func setupApp() *gofr.App {
	app := gofr.New()

	// Set up a mock database for testing
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	app.SetDB(db)

	// Set up routes as in the main function

	return app
}

// Utility function to insert test data into the database
func insertTestData(db *sql.DB) {
	_, err := db.Exec(`
		INSERT INTO Student_details (first_name, last_name, age, class, gender, address, phone_number) 
		VALUES ('John', 'Doe', 25, 1, 'Male', 'Address', 1234567890),
			   ('Jane', 'Doe', 30, 2, 'Female', 'AnotherAddress', 9876543210)
	`)
	if err != nil {
		panic(err)
	}
}

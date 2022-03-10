package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"strconv"
)

var (
	db  *sql.DB
	err error
)

// checkErr: handling errors
func checkErr(err error) {
	if err != nil {
		fmt.Println("Error found")
		panic(err)
	}
}

// printMessage: handling messages
func printMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

// DB set up init
func init() {
	connStr := "postgres://postgres:password@localhost/college?sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	checkErr(err)
}

type Student struct {
	Name       string `json:"name"`
	Department string `json:"department"`
	Address    string `json:"address"`
}

type JsonResponse struct {
	Type    string    `json:"type"`
	Data    []Student `json:"data"`
	Message string    `json:"message"`
}

func main() {
	// Init the mux router
	router := mux.NewRouter()

	// Get all Students
	router.HandleFunc("/student", GetStudents).Methods("GET")

	// Create a Student
	router.HandleFunc("/student", CreateStudent).Methods("POST")

	// Update a Student
	router.HandleFunc("/student", UpdateStudent).Methods("PUT")

	// Delete a specific Student by the StudentID
	router.HandleFunc("/student/{id}", DeleteStudent).Methods("DELETE")

	// Delete all Students
	router.HandleFunc("/student", DeleteStudents).Methods("DELETE")

	// serve the app at port 8000
	fmt.Println("Server at 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

// Get all Students
func GetStudents(w http.ResponseWriter, r *http.Request) {

	printMessage("Getting Students...")

	// Get all Students from Students table
	rows, err := db.Query("SELECT * FROM students")
	defer rows.Close()

	// check errors
	checkErr(err)

	// var response []JsonResponse
	var Students []Student

	// Foreach Student
	for rows.Next() {
		var id int
		var StudentName string
		var StudentDepartment string
		var StudentAddress string

		err = rows.Scan(&id, &StudentName, &StudentDepartment, &StudentAddress)

		// check errors
		checkErr(err)

		Students = append(Students, Student{Name: StudentName, Department: StudentDepartment, Address: StudentAddress})
	}

	var response = JsonResponse{Type: "success", Data: Students}

	json.NewEncoder(w).Encode(response)
}

// Create a Student
func CreateStudent(w http.ResponseWriter, r *http.Request) {
	StudentName := r.FormValue("name")
	StudentDepartment := r.FormValue("department")
	StudentAddress := r.FormValue("address")

	var response = JsonResponse{}

	if StudentName == "" && StudentDepartment == "" && StudentAddress == "" {
		response = JsonResponse{Type: "error", Message: "Please provide name and department and address parameter."}
	} else {

		printMessage("Inserting Student into DB")

		fmt.Println("Inserting new Student with name: " + StudentName + " and department: " + StudentDepartment + " and Student Address: " + StudentAddress)

		var lastInsertID int
		err := db.QueryRow("INSERT INTO students(name, department, address) VALUES($1, $2, $3) returning id;", StudentName, StudentDepartment, StudentAddress).Scan(&lastInsertID)

		// check errors
		checkErr(err)

		response = JsonResponse{Type: "success", Message: fmt.Sprintf("The Student id %d has been inserted successfully!", lastInsertID)}
	}

	json.NewEncoder(w).Encode(response)
}

// Update a Student
func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	StrStudentId := r.FormValue("id")
	StudentId, err := strconv.Atoi(StrStudentId)
	checkErr(err)
	StudentDepartment := r.FormValue("department")
	StudentAddress := r.FormValue("address")

	response := JsonResponse{}

	if StudentId == 0 && StudentDepartment == "" && StudentAddress == "" {
		response = JsonResponse{Type: "error", Message: "Please provide id, department and address parameter."}
	} else {
		_, err := db.Exec("UPDATE Students SET department = $2, address = $3 where id = $1", StudentId, StudentDepartment, StudentAddress)

		// check errors
		checkErr(err)

		printMessage("Updating Student into DB")

		details := fmt.Sprintf("Updating Student with id: %d and department: %s and  address: %s", StudentId, StudentDepartment, StudentAddress)
		fmt.Println(details)

		response = JsonResponse{Type: "success", Message: fmt.Sprintf("The StudentID=%d has been updated successfully!", StudentId)}
	}
	json.NewEncoder(w).Encode(response)
}

// Delete a Student
func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	StrStudentID := params["id"]
	StudentID, err := strconv.Atoi(StrStudentID)
	checkErr(err)

	var response = JsonResponse{}

	if StudentID == 0 {
		response = JsonResponse{Type: "error", Message: "You are missing id parameter."}
	} else {

		printMessage("Deleting Student from DB")

		_, err := db.Exec("DELETE FROM Students where id = $1", StudentID)

		// check errors
		checkErr(err)

		response = JsonResponse{Type: "success", Message: fmt.Sprintf("The StudentID=%d has been deleted successfully!", StudentID)}
	}

	json.NewEncoder(w).Encode(response)
}

// Delete all Students
func DeleteStudents(w http.ResponseWriter, r *http.Request) {

	printMessage("Deleting all Students...")

	_, err := db.Exec("DELETE FROM Students")

	// check errors
	checkErr(err)

	printMessage("All Students have been deleted successfully!")

	var response = JsonResponse{Type: "success", Message: "All Students have been deleted successfully!"}

	json.NewEncoder(w).Encode(response)
}

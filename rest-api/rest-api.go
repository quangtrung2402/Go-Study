package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"

	"strconv"
)

type Course struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

var courses []Course

func main() {
	r := mux.NewRouter()
	courses = append(courses, Course{
		ID:    "1",
		Title: "The first course",
		Body:  "The first course content"})

	courses = append(courses, Course{
		ID:    "2",
		Title: "The second course",
		Body:  "The second course content"})

	// Routes consist of a path and a handler function.
	r.HandleFunc("/courses", createCourse).Methods("POST")
	r.HandleFunc("/courses/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/courses", getCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", getCourse).Methods("GET")
	r.HandleFunc("/courses/{id}", deleteCourse).Methods("DELETE")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}

/** Postman test
Method: GET
Request URL: localhost:8000/courses
*/
func getCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

/** Postman test
Method: GET
Request URL: localhost:8000/courses/2
*/
func getCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	for _, item := range courses {
		if item.ID == param["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Course{})
}

/** Postman test
Method: POST
Request URL: localhost:8000/courses
Body: raw
	{
        "title": "New course",
        "body": "This is content of new course"
    }
*/
func createCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	course.ID = strconv.Itoa(rand.Intn(1000000))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(&course)
}

/** Postman test
Method: PUT
Request URL: localhost:8000/courses/1
Body: raw
	{
        "title": "Updated first course",
        "body": "Updated content of first course"
    }
*/
func updateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range courses {
		if item.ID == params["id"] {
			courses = append(courses[:i], courses[i+1:]...)

			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.ID = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(&course)
			return
		}
	}
	json.NewEncoder(w).Encode(courses)
}

/** Postman test
Method: DELETE
Request URL: localhost:8000/courses/1
*/
func deleteCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range courses {
		if item.ID == params["id"] {
			courses = append(courses[:i], courses[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(courses)
}

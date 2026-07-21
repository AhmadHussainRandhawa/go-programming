package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string `json:"courseid"`
	CourseName  string `json:"coursename"`
	CoursePrice int    `json:"price"`
	Author      *Author
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// middleeware, helper
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""

}

// fake db
var courses []Course

func main() {

}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to my API</h1>"))

}

func getAllCourses(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Get all courses")

	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w) // reates an encoder that knows: "When I produce JSON, I'll send it to w.
	encoder.Encode(courses)

}

func getOneCourse(w http.ResponseWriter, r *http.Request) {

	// grab id from request
	params := mux.Vars(r)

	// loop through courses, find matching id and return the response.
	for i, course := range courses {

		fmt.Println(i, course.CourseId, course.CourseName)

		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	encoder := json.NewEncoder(w)
	encoder.Encode("No course found with give id")
	return

}

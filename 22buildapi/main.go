package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

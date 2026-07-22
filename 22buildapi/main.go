package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

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
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""

}

func main() {

	r := mux.NewRouter()

	// Seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "Java",
		CoursePrice: 199, Author: &Author{FullName: "Ahmad Hussain", Website: "ahmad.dev"}})

	courses = append(courses, Course{CourseId: "3", CourseName: "python",
		CoursePrice: 299, Author: &Author{FullName: "Farman Ali", Website: "farmantechmaster.dev"}})

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe("127.0.0.1:8000", r))

}

// fake db
var courses []Course

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to my API</h1>"))

}

func getAllCourses(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Get all courses")

	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w) // Creates an encoder that knows: "When I produce JSON, I'll send it to w.
	encoder.Encode(courses)       // Convert the Go value stored in course (or courses) into JSON and write that JSON to the destination (w).

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

}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	var course Course
	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Please send valid JSON data."})
		return
	}

	for _, c := range courses {
		if course.CourseName == c.CourseName {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "A course with this name is already exists"})
		}
	}

	if course.IsEmpty() {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Course name is required."})
		return
	}

	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(course)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	params := mux.Vars(r)

	var updatedCourse Course
	err := json.NewDecoder(r.Body).Decode(&updatedCourse)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid or missing JSON body."})
		return
	}

	for index, course := range courses {
		if course.CourseId == params["id"] {
			updatedCourse.CourseId = params["id"] // preserve original ID
			courses[index] = updatedCourse        // in-place replace, no delete needed
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(updatedCourse)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Course with given id not found."})
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicatioan/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			// TODO: send a confirm or deny response
			break
		}
	}
}

// Encoder: Go → JSON → Response
// Decoder: Request → JSON → Go

package main

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

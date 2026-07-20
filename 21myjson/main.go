package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"courseName"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              // Do not expose this field in the JSON output.
	Tags     []string `json:"tags,omitempty"` // Omit this field from the JSON output if it is empty or nil.
}

func main() {

	EncodeJson()

}

func EncodeJson() {

	udemyCourses := []course{
		{"Python Bootcamp", 199, "udemy.com", "abc123", []string{"ML", "AI", "DL"}},
		{"javascript Bootcamp", 299, "udemy.com", "cew123", []string{"web-dev", "js"}},
		{"Mern Stack", 99, "udemy.com", "tri123", nil},
	}

	// Package this data as JSON data.

	dataByte, _ := json.MarshalIndent(udemyCourses, "", "\t")

	fmt.Println(string(dataByte))

}

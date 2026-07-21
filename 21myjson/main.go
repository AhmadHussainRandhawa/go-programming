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

	// EncodeJson()
	// DecodeJson()

}

func EncodeJson() {

	udemyCourses := []course{
		{"Python Bootcamp", 199, "udemy.com", "abc123", []string{"ML", "AI", "DL"}},
		{"javascript Bootcamp", 299, "udemy.com", "cew123", []string{"web-dev", "js"}},
		{"Mern Stack", 99, "udemy.com", "tri123", nil},
	}

	// Package this data as JSON data.

	dataByte, _ := json.MarshalIndent(udemyCourses, "", "    ")

	fmt.Println(string(dataByte))

}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
		{
			"courseName": "Python Bootcamp",
			"Price": 199,
			"website": "udemy.com",
			"tags": ["ML", "AI", "DL"]
		}
	`)

	checkJsonValid := json.Valid(jsonDataFromWeb)

	var udemyCourses course

	if checkJsonValid {

		err := json.Unmarshal(jsonDataFromWeb, &udemyCourses)

		if err != nil {
			fmt.Println(err)
			return
		}

	} else {
		fmt.Println("This is not valid JSON")
	}

	fmt.Printf("%#v", udemyCourses)

	// Some cases where you just want to store json data to a key, value pair.

	var mapOnlineData map[string]any
	err := json.Unmarshal(jsonDataFromWeb, &mapOnlineData)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("\n\n%#v", mapOnlineData)

	for key, value := range mapOnlineData {
		fmt.Printf("\nKey is %v and value is %v", key, value)
	}

}

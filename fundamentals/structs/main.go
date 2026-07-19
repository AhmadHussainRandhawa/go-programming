package main

import "fmt"

func main() {
	fmt.Println("Welcome, structs in golang")

	// No inheritence in golang; no super, no parent.

	ahmad := User{"Ahmad", "ahmad@go.dev", true, 21}

	fmt.Println(ahmad)
	fmt.Printf("Ahmad Details are: %+v\n", ahmad)
	fmt.Printf("Ahmad age is: %+v\n", ahmad.age)

}

type User struct {
	Name   string
	Email  string
	status bool
	age    int
}

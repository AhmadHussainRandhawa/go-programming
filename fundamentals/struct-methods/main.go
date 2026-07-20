package main

import "fmt"

func main() {
	ahmad := User{"Ahmad", "ahmad@go.dev", true, 21}
	fmt.Println(ahmad)
	fmt.Printf("Ahmad Details are: %+v\n", ahmad)
	fmt.Printf("Ahmad age is: %v\n", ahmad.Age)

	ahmad.Greet()

	ahmad.newEmail()
	fmt.Println("Email remains unchanged:", ahmad.Email)

	rect := Rectangle{Width: 4, Height: 5}
	fmt.Println("Area:", rect.Area())
	fmt.Println("Perimeter:", rect.Perimeter())
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) Greet() {
	fmt.Println("Welcome", u.Name)
}

func (u User) newEmail() {
	u.Email = "jutt@py.com"
	fmt.Println("New Email:", u.Email)
}

type Rectangle struct {
	Width  int
	Height int
}

func (r Rectangle) Area() int {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() int {
	return 2 * (r.Width + r.Height)
}

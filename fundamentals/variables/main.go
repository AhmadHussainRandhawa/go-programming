package main

import "fmt"

const LoginToken string = "example-login-token" // Public

func main() {
	var name string = "Ahmad"
	fmt.Println(name)
	fmt.Printf("Variable is of type: %T\n", name)

	var IsMale bool = true
	fmt.Println(IsMale)
	fmt.Printf("Variable is of type: %T\n", IsMale)

	var salary uint8 = 255
	fmt.Println(salary)
	fmt.Printf("Variable is of type: %T\n", salary)

	var height float32 = 59.32235252390540900
	fmt.Println(height)
	fmt.Printf("Variable is of type: %T\n", height)

	// default value and some alias
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T\n", anotherVariable)

	// implicit type ... Most common (outside function)

	var journal = "whtatljs icmas sal"
	fmt.Println(journal)

	// no var style  ... Most common (inside function)

	numberOfUser := 34200.93
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T\n", LoginToken)

}

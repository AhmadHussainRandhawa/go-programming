package main

import "fmt"

func main() {
	var ptr *int

	fmt.Println("the default value of my pointer is:", ptr)

	myNumber := 11
	ptr = &myNumber

	fmt.Println("the ptr value:", ptr)
	fmt.Println("the ptr value:", *ptr)

	*ptr = *ptr * 2
	*ptr = *ptr - 3

	fmt.Println("New value of ptr:", *ptr)
	fmt.Println("variable myNumber value is:", myNumber)

}

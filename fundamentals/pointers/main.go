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

	// pointer with structs

	p := &Person{
		Name: "Ahmad",
		Age:  20,
	}

	fmt.Println(p)

	// Interesting flow

	x := 10
	ptr = &x

	s := Student{"Ali", 21}
	sptr := &s

	fmt.Println(ptr)
	fmt.Println(sptr)

	fmt.Printf("%p\n", ptr)
	fmt.Printf("%p\n", sptr)

}

type Person struct {
	Name string
	Age  int
}

type Student struct {
	Name string
	Age  int
}

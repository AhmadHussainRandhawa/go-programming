package main

import "fmt"

func main() {
	fmt.Println("World")

	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Hello")
	fmt.Println("Three")

	myDefer()

}

// one, two, hello
// 0, 1, 2, 3, 4
//
// output: World, Three, line,  43210Hello, Two, One

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
	fmt.Print("\n")
}

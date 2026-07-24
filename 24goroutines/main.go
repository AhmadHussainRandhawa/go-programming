package main

import (
	"fmt"
	"time"
)

func test(val string) {
	fmt.Println(val)
}

func main() {
	go test("Hello")
	go test("Hello2")

	time.Sleep(time.Second * 3)

	fmt.Println("world")
}

// Goroutines — create concurrent execution.

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup // WaitGroup — wait for concurrent work to finish.

	wg.Add(3)

	go test(1, &wg)
	go test(2, &wg)
	go test(3, &wg)

	fmt.Println("This is main before waiting...")

	wg.Wait()

	fmt.Println("This is main after waiting...")
}

func test(id int, wg *sync.WaitGroup) {
	fmt.Println("worker...", id)

	defer wg.Done()

}

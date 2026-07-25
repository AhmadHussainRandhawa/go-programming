package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter int

var mu sync.Mutex

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go increment(&wg)
	go increment(&wg)

	wg.Wait()

	fmt.Println("counter with mutex:", counter)

}

// Problem - Predict the output

// func increment(wg *sync.WaitGroup) {

// 	defer wg.Done()

// 	temp := counter   // Read
// 	runtime.Gosched() // Yield to another goroutine - It tells the scheduler: "I'm willing to pause here. Let another runnable goroutine execute."
// 	temp++            // Modify
// 	counter = temp    // Write

// }

// Solution - Mutex (lock)

func increment(wg *sync.WaitGroup) {

	defer wg.Done()

	mu.Lock()
	defer mu.Unlock()

	temp := counter
	runtime.Gosched()
	temp++
	counter = temp

}

// `go run --race main.go` - to check for race condition.

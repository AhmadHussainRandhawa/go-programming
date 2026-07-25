package main

import (
	"fmt"
	"time"
)

// func main() {

// 	ch := make(chan int)

// 	go func() {
// 		ch <- 21
// 	}()

// 	val := <-ch
// 	fmt.Println(val)

// }

func main() {
	numChan := make(chan int)

	go processNum(numChan)
	numChan <- 10

	time.Sleep(2 * time.Second)
}

func processNum(numChan chan int) {
	fmt.Println("Processing the number:", <-numChan)
}

package main

import (
	"errors"
	"fmt"
)

func main() {

	greet("Ahmad Hussain")

	resultAdd := add(4, 5)
	fmt.Println(resultAdd)

	a, b := swap(10, 33)
	fmt.Println("Swap values are:", a, b)

	resultDiv, err := divide(9, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Result:", resultDiv)
	}

	fmt.Println(sum(2, 3, 5, 10))

	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(numbers...))

}

func greet(name string) {
	fmt.Println("Hello", name)
}

func add(a int, b int) int {
	return a + b
}

func swap(x, y int) (int, int) {
	return y, x
}

func divide(a, b float32) (float32, error) {
	if b == 0 {
		return 0, errors.New("cannot be divisible by 0")
	}

	return a / b, nil
}

func sum(nums ...int) int {
	total := 0
	for _, value := range nums {
		total += value
	}
	return total
}

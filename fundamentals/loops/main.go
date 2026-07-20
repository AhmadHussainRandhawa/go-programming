package main

import "fmt"

func main() {
	days := []string{"Friday", "Saturday", "Sunday", "Tuesday", "Wednesday", "Thursday"}

	// 1. Basic for Loop
	for i := 0; i < len(days); i++ {
		fmt.Println("Today is:", days[i])
	}

	// 2. Looping with range (index)
	for i := range days {
		fmt.Println("Another today is:", days[i])
	}

	// 3. Looping with range (index, value)
	for i, j := range days {
		fmt.Printf("Today (%v) is %v\n", i, j)
	}

	// 4. Interesting.

	x := 1

	for x < 10 {

		if x == 3 {
			x++
			continue

		} else if x == 6 {
			goto ahr
		}

		fmt.Println("This is x: ", x)
		x++
	}

ahr:
	fmt.Println("Redirectring to website: ahmadhussainrandhawa.com")
}

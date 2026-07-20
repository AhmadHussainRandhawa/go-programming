package main

import "fmt"

func main() {
	age := 20
	hasID := true
	marks := 82

	// 1. if-else
	if hasID {
		fmt.Println("You have an ID card.")
	} else {
		fmt.Println("You don't have an ID card.")
	}

	// 2. if-else if-else chain
	if marks >= 90 {
		fmt.Println("Grade: A+")
	} else if marks >= 80 {
		fmt.Println("Grade: A")
	} else if marks >= 70 {
		fmt.Println("Grade: B")
	} else {
		fmt.Println("Grade: C")
	}

	// 3. Multiple conditions (&& and ||)
	if age >= 18 && hasID {
		fmt.Println("You can enter the event.")
	}

	if age < 18 || !hasID {
		fmt.Println("Entry denied.")
	} else {
		fmt.Println("All requirements satisfied.")
	}

	// 4. Nested if
	if age >= 18 {
		if hasID {
			fmt.Println("Nested: Adult with ID.")
		} else {
			fmt.Println("Nested: Adult without ID.")
		}
	}

	// 5. Short statement before if
	if score := 95; score >= 90 {
		fmt.Println("Excellent score:", score)
	} else {
		fmt.Println("Score:", score)
	}

	// score doesn't exist here
	// fmt.Println(score) // Compilation error

	if num := 10; num > 10 {
		fmt.Println("Greater then 10")
	} else {
		fmt.Println("Not greater then 10")
	}
}

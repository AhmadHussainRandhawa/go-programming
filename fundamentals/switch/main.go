package main

import (
	"fmt"
	"math/rand"
)

func main() {
	diceNumber := rand.Intn(8) + 1
	fmt.Println("value of dice is:", diceNumber)

	switch diceNumber {
	case 0, 1:
		fmt.Println("You can move one spot")

	case 2:
		fmt.Println("You can move two spot")

	case 3:
		fmt.Println("You can move three spot")

	case 4:
		fmt.Println("You can move four spot")
		fallthrough

	case 5:
		fmt.Println("You can move five spot")

	case 6:
		fmt.Println("You can move six spot")

	default:
		fmt.Println("What wass that!")
	}
}

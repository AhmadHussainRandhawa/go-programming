package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{"Apple", "Mango", "Banana"}
	fmt.Printf("Type of fruit list is: %T\n", fruitList)

	fruitList = append(fruitList, "Peach", "Guava")
	fmt.Println("Fruit List:", fruitList)

	newList := append(fruitList[:3])
	fmt.Println("New List:", newList)

	// make(data_type, length)
	highScores := make([]int, 3)

	highScores[0] = 234
	highScores[1] = 256
	highScores[2] = 984
	// highScores[3] = 482   // Un-comment it, You can not add values like that.

	fmt.Println(highScores)

	highScores = append(highScores, 5120, 8974350, 32, 2, 908)

	fmt.Println(highScores)

	sort.Ints(highScores)

	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	// How to remove a value from slices based on index

	courses := []string{"Django", "Linux Internals", "Git/Github", "Databases", "SQL", "System Design", "Networking"}
	var index int = 2

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}

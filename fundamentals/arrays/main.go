package main

import "fmt"

func main() {
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Mango"

	fmt.Println("Fruit list includes:", fruitList)

	var vegList = [5]string{"Potato", "Mushrooms", "Beans"}

	fmt.Println("Vegetables list includes:", vegList)
	fmt.Println("Length of Vegetables list:", len(vegList))
}

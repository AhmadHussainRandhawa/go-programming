package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating of my pizza:")

	input, _ := reader.ReadString('\n')
	fmt.Println("The rating is:", input)
}

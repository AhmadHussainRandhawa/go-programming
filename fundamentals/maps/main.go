package main

import "fmt"

func main() {
	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Rubby"
	languages["PY"] = "Python"
	languages["RS"] = "Rust"

	fmt.Println("Lists of all languages:", languages)
	fmt.Println("PY shorts for:", languages["PY"])

	// loops are interesting in GO

	for _, value := range languages {
		fmt.Printf("For key v, value is %v\n", value)
	}

}

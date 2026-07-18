package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://www.york.ac.uk/teaching/cws/wws/webpage1.html"

func main() {
	response, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer response.Body.Close()
	fmt.Printf("Response is of type: %T\n", response)

	databyte, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Response body is:,", string(databyte))

}

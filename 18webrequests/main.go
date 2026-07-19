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
	fmt.Println("\n\n", response)
	fmt.Println("\n\n", *response)
	fmt.Printf("\n\n%p\n\n", response)

	databyte, err := io.ReadAll(response.Body) // internally go does: *response.Body... You can write: (*response).Body

	fmt.Printf("%T\n\n", response.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Response body is:,", string(databyte))

}

func PerformGetRequest(myurl string) {
	response, err := http.Get(myurl)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer response.Body.Close()

	fmt.Println(response.StatusCode)
	fmt.Println(response.ContentLength)
	fmt.Println(response.Body)

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}

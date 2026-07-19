package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	PerformGetRequest("http://localhost:8000")

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

	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount is: ", byteCount)
	fmt.Println(responseString.String())

	// fmt.Println(string(content))

}

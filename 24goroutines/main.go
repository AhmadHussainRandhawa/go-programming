package main

import (
	"fmt"
	"net/http"
)

func getStatusCode(endpoint string) {

	resp, err := http.Get(endpoint)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%d is status code for website: %s\n", resp.StatusCode, endpoint)

}

func main() {
	websitelist := []string{
		"https://www.google.com/",
		"https://www.youtube.com/",
		"https://www.github.com/",
		"https://www.linkedin.com/",
	}

	for _, web := range websitelist {
		getStatusCode(web)
	}

}

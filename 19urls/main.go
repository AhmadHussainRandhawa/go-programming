package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://pkg.go.dev:8000/net/url?tab=doc&rough=test#URL"

func main() {
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Hostname())
	fmt.Println(result.Port())
	fmt.Println(result.Path)

	fmt.Println(result.RawQuery)
	fmt.Println(result.Query())

	qparams := result.Query()

	fmt.Println(qparams["tab"])

	for _, value := range qparams {
		fmt.Println("Result query values are:", value)
	}

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "pkg.go.dev",
		Path:     "/net/url",
		RawQuery: "tab=doc",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)

}

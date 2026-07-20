package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// PerformGetRequest("http://localhost:8000")
	// PerformPostJsonRequest("http://localhost:8000/post")
	PerformPostFormRequest("http://localhost:8000/postform")

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

func PerformPostJsonRequest(myurl string) {

	requestBody := strings.NewReader(`
		{
			"courseName": "Go-Programming", 
			"userName": "Ahmad Hussain",
			"price": 23599

		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer response.Body.Close()

	dataByte, _ := io.ReadAll(response.Body)

	fmt.Println("Response content:", string(dataByte))

}

func PerformPostFormRequest(myurl string) {

	data := url.Values{}

	data.Add("firstName", "ahmad")
	data.Add("lastName", "hussain")
	data.Add("identity", "Muslim")

	response, _ := http.PostForm(myurl, data)

	contentByte, _ := io.ReadAll(response.Body)

	defer response.Body.Close()

	fmt.Printf(string(contentByte))
}

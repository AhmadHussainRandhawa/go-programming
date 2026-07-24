package main

import (
	"fmt"
	"net/http"

	"github.com/AhmadHussainRandhawa/mongoapi/router"
)

func main() {
	fmt.Println("Server is listening...")

	r := router.Router()
	http.ListenAndServe(":4000", r)

}

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Print("Welcome to my time\n")

	presentTime := time.Now()

	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 Monday 15:04:05"))

	fmt.Println(presentTime.Format("Jan 02 15:04"))

	fmt.Println(presentTime.Format("Jan 02 03:04 PM"))

	// Create your own date
	created_date := time.Date(2005, time.May, 26, 17, 0, 0, 0, time.UTC)
	fmt.Println(created_date)

	fmt.Println(created_date.Format("Mon 03:04 PM Jan-2006"))

}

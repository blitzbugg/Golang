package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)
	
	fmt.Println(presentTime.Format("01-02-2006 Monday"))

	createdDate := time.Date(2003, time.December, 21, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
	
}

package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers in go")

	// var ptr *int 
	// fmt.Println("The value of ptr is :", ptr)

	myNumber := 23

	var ptr = &myNumber

	fmt.Println("Value of ptr mynumber is :", ptr)
	fmt.Println("Value of ptr mynumber is :", *ptr)

	*ptr = *ptr * 2
	fmt.Println("New value is :", myNumber)
}

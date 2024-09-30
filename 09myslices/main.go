package main

import "fmt"

func main() {
	fmt.Println("Welcome to Slices")

	var fruitList = []string{"Apple", "Orange", "Grape"}
	fmt.Printf("Type of fruit list is: %T", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Printf("\nFruit list after appending: %v\n", fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)
}

package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Cherry"

	fmt.Println("Fruit list is :", fruitList)
	fmt.Println("Fruit list length is :", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("The veglist is :", vegList)
}

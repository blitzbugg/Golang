package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices")

	var fruitList = []string{"Apple", "Orange", "Grape"}
	fmt.Printf("Type of fruit list is: %T", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Printf("\nFruit list after appending: %v\n", fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 123
	highScores[1] = 456
	highScores[2] = 789
	highScores[3] = 109

	highScores = append(highScores, 443, 343, 432)

	fmt.Println(highScores)

	fmt.Print(sort.IntsAreSorted(highScores))


	sort.Ints(highScores)

	fmt.Println("Sorted: ",highScores)

	fmt.Print(sort.IntsAreSorted(highScores))

	// How to remove a value from slices based on index

	var courses = []string{"reactjs", "Javascript", "Python", "Django"}
	fmt.Println("Before removing: ", courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("After removing: ", courses)

}

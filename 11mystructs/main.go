package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// No inheritance in golang; No super or parent

	blitzbugg := User{"blitzbugg", "blitzbugg@go.dev", true, 20}
	fmt.Println(blitzbugg)
	fmt.Printf("blitzbugg details are: %+v\n", blitzbugg)
	fmt.Printf("Name is :%v and his age is %v", blitzbugg.Name, blitzbugg.Age)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

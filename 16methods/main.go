package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// No inheritance in golang; No super or parent

	blitzbugg := User{"blitzbugg", "blitzbugg@go.dev", true, 20}
	fmt.Println(blitzbugg)
	fmt.Printf("blitzbugg details are: %+v\n", blitzbugg)
	fmt.Printf("Name is :%v and his age is %v\n", blitzbugg.Name, blitzbugg.Age)
	blitzbugg.GetStatus()
	blitzbugg.NewMail()
	fmt.Println("Email of blitzbugg: ",blitzbugg.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int

}

func (u User) GetStatus() {
	fmt.Println("Is user is active: ", u.Status)
}

func (u User) NewMail(){
	u.Email = "abc@gmail.com"
	fmt.Println("Email of the user is :", u.Email)
}
package main

import "fmt"

func main() {
	fmt.Println("Welcome to oops in golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days{
	// 	fmt.Printf("The index is %v and days is %v\n",index+1,day)
	// }

	rougueValue := 1

	for rougueValue < 10 {

		if rougueValue == 2 {
			goto label
		}

		if rougueValue == 5 {
			break
		}
		fmt.Println("Value is ", rougueValue)
		rougueValue++
	}

	// Goto
	label:
		fmt.Println("Just landed here")
}
package main

import "fmt"

func main() {
	fmt.Println("function in golang")
	greeter()
	
	result := adder(3,4)
	fmt.Println("The result is: ",result)

	proResult := proAdder(2,5,6,7,8,3)
	fmt.Println("The pro result is: ",proResult)
}

func greeter()  {
	fmt.Println("Namaskkaram :)")
}

func proAdder(values ...int) int {
	total := 0

	for _, val := range values{
		total += val
	}
	return total
}

func adder(num1 int, num2 int) int {
	return num1 + num2
}
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files in golang")
	content := "This is a sample content"

	file, err := os.Create("./myfile.txt")
	checkNilErr(err)

	// if err != nil {
	// 	panic(err)
	// }

	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("length is: ", length)
	defer file.Close()
	readFile("./myfile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)

	fmt.Println("Text data inside the file is \n", string(databyte))
}


func checkNilErr(err error)  {
	if err != nil {
		panic(err)
	}
}
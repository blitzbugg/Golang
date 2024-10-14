package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Web verb")
	PerformGetRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:3000/get"

	resp, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println("Status code: ", resp.StatusCode)
	fmt.Println("Content length: ", resp.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(resp.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("Byte count is: ", byteCount)
	fmt.Println(responseString.String())
	
	// fmt.Println(content)
	// fmt.Println(string(content))
}

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Web verb")
	// PerformGetRequest()
	performPostJsonRequest()
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

func performPostJsonRequest()  {
	const myurl = "http://localhost:3000/post"

	// Fake json playload

	requestBody := strings.NewReader(`
		{
	"coursename": "ReactJS",
	"price": "0",
	"platform": "react.dev"
	}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}
package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://www.youtube.com/watch?v=Q3uQctGscg0"

func main() {
	fmt.Println("Handling urls in golang")
	fmt.Println(myUrl)

	// parsing
	result, _ := url.Parse(myUrl)

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)

	fmt.Println(qparams["v"])

	// for _, val := range qparams {
	// 	fmt.Println("Param is: ",val)
	// }

	// There is only one param in the url that I have provided so There is no need to use the loop

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "www.youtube.com",
		Path:     "/watch",
		RawQuery: "v=Q3uQctGscg0",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
} 

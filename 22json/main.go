package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name string `json:"coursename"`
	Price int 
	Platform string `json:"website"`
	Password string `json:"-"`
	Tags []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("How to make json in golang")
	// EncodeJson()
	DecodeJson()
}


func EncodeJson()  {
	
	Courses := []course{
		{"Go", 100, "Udemy", "abc123", []string{"web-dev", "Go"}},
		{"Python", 200, "Udemy", "abc123", []string{"web-dev", "python"}},
		{"Kotlin", 300, "Udemy", "abc123", nil},
	}

	// Package this data a JSON data

	finalJson, err := json.MarshalIndent(Courses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson()  {
	jsonDataFromWeb := []byte(`
	 {
         "coursename": "Go",
         "Price": 100,
         "website": "Udemy",
         "tags": ["web-dev","Go"]
    }
	`)

	var udemyCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &udemyCourse)
		fmt.Printf("%#v\n", udemyCourse)
	} else {
		fmt.Println("JSON was not valid")
	}

	// Some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData{
		fmt.Printf("\nKey is %v and value is %v and type is: %T", k, v, v)
	}

}
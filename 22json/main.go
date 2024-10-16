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
	EncodeJson()
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
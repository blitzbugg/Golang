package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Model for course - file
type Course struct {
	COurseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// Fake DB
var courses []Course

// middleware, helper -file
func isEmpty(c *Course) bool {
	return c.COurseId == "" && c.CourseName == ""
}
func main() {

}

// Controllers - file

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome go API</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-type", "application/json")

	// Grab id from request
	params := mux.Vars(r)
	fmt.Printf("Type of %v is %T", params, params)

	for _, course := range courses {
		if course.COurseId == params["id"] {
			json.NewEncoder(w).Encode(course.Author)
			return
		}
		json.NewEncoder(w).Encode("No course id with give id")
		return
	}
}

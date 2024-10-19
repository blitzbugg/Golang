package main

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

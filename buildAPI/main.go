package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Model for course -file
type Course struct {
	CourseId    string  `json:"courseId"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware, helper -file
func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

func main() {
	fmt.Println("API")
	r := mux.NewRouter()
	//seeding
	courses = append(courses, Course{
		CourseId:    "2",
		CourseName:  "Reactjs course",
		CoursePrice: 399,
		Author: &Author{
			Fullname: "Karanveer Sharma",
			Website:  "karanveer-portfolio.vercel.app",
		},
	})
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")
	//listening to port
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}

//controllers -file

//serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println(courses)
	w.Write([]byte("<h1>Welcome to API by Karanveer Sharma</h1>"))
}
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")
	//grab id from request

	params := mux.Vars(r)

	// loop through courses, find matching id and return response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send somee data")
		return
	}
	var course Course
	_ = json.NewDecoder((r.Body)).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside json")
		return
	}
	//generate unique id,string
	//append course into courses
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
}
func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Deletion successful")
			return
		}
	}
	json.NewEncoder(w).Encode("Id not found")
}

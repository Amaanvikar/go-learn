package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// model for my course - file
type Course struct {
	CourseName  string  `json:"coursename"`
	CourseId    string  `json:"courseid"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// dummy DB
var courses []Course

// middleware, helper - file
func (c *Course) IsEmpty() bool {
	// return c.CourseName == "" && c.CourseId == ""
	return c.CourseName == ""
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")

	fmt.Println("Server running on :8000")
	http.ListenAndServe(":8000", r)
}

// controllers - file

// serve home  route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>This is Go-lang series</h1>"))
}

func getAllCourses(w http.ResponseWriter, r http.Request) {
	fmt.Println("Get All courses")
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("content:type", "application/json")

	// grab id from request

	params := mux.Vars(r)
	fmt.Println("params", params)
	fmt.Printf("The type of param: %T\n", params)

	// loop through courses, find matching id and return the response

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course found with the given id")
	return

}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Course created successfully")
	w.Header().Set("content:type", "application/json")

	//  what if the body is empty

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// what about - {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("NO data inside json")
		return
	}

	// generate unique id, string
	// append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Course updated Successfully")
	w.Header().Set("Content-Type", "application/json")

	// first grab id from req

	params := mux.Vars(r)

	// loo , id , remove , add with my ID

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

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}

}

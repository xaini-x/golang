package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice string  `json:"courseprice"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// Middleware function
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

// controllers for testing
func ServerHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome to api </h1>"))
}
func main() {
	fmt.Println("hello world")
	r := mux.NewRouter()

	// Create a new database
	courses = append(courses, Course{CourseId: "111", CourseName: "Course1", CoursePrice: "123", Author: &Author{Fullname: "xaini 's", Website: "github.com/xaini-x"}})
	courses = append(courses, Course{CourseId: "111", CourseName: "Course2", CoursePrice: "167", Author: &Author{Fullname: "xerx", Website: "github.com/f"}})
	courses = append(courses, Course{CourseId: "111", CourseName: "Course3", CoursePrice: "133", Author: &Author{Fullname: "zoro", Website: "github.com/g"}})
	courses = append(courses, Course{CourseId: "111", CourseName: "Course4", CoursePrice: "138", Author: &Author{Fullname: "aryaa", Website: "github.com/v"}})
	//listen to port

	r.HandleFunc("/", ServerHome).Methods("GET")
	r.HandleFunc("/getCourses", GetAllCourses).Methods("GET")
	r.HandleFunc("/oneCourse/{id}", GetOneCourse).Methods("GET")
	r.HandleFunc("/create", CreateOneCourse).Methods("POST")
	r.HandleFunc("/updateOne/{id}", UpdateOnrCourse).Methods("PUT")
	r.HandleFunc("/deleteOne/{id}", deleteCourse).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("all COurses")
	w.Header().Set("content-type", "encoding/json")
	json.NewEncoder(w).Encode(courses)
}

func GetOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("detail of ur course")
	w.Header().Set("Content-type", "Encoding/json")
	//grab id from request
	params := mux.Vars(r)
	// grab param and search in db
	for _, course := range courses {
		//check if course exists
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("no course found")
	return
}

func CreateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create a new course")
	w.Header().Set("Content-Type", "application/json")
	// if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("no data found")
	}
	//{} json data found
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("no data inside the json body")
		return
	}
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(courses)
	return
}

func UpdateOnrCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Print("updateone course")
	w.Header().Set("Content-type", "application/json")
	// GRAB ID FROM PARAMS
	params := mux.Vars(r)

	// loop thorough index to check if exist
	for index, course := range courses {
		//check if course exists
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(courses)
			return
		}
	}
	//if id not found return
	json.NewEncoder(w).Encode("no course found")
	return

}

func deleteCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Print("delete course")
	w.Header().Set("Content-type", "application/json")
	// GRAB ID FROM PARAMS
	params := mux.Vars(r)

	// loop thorough index to check if exist
	for index, course := range courses {
		//check if course exists
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode("course deleted ")
	return

}

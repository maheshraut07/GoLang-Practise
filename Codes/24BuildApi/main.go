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

// Course model
type Course struct {
	CourseId   string  `json:"courseid"`
	CourseName string  `json:"coursename"`
	CoursePrice int     `json:"-"`
	Author     *Author `json:"author"`
}

// Author model
type Author struct {
	FirstName string `json:"fullname"`
	Website   string `json:"website"`
}

// Fake database
var courses []Course  // courses is the slice of the type course 

func main() {
	fmt.Println("Welcome to the lecture of the building of API")

	// Create router using Gorilla Mux
	r := mux.NewRouter()

	// Seed fake data , append the values to the courses slice 
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJs", CoursePrice: 299, Author: &Author{FirstName: "mahesh raut ", Website: "Lco.dev"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "MERN stack", CoursePrice: 199, Author: &Author{FirstName: "mahesh raut ", Website: "Lco.dev"}})

	// Routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// Listen on port 4000
	fmt.Println("Starting server on port 4000...")
	log.Fatal(http.ListenAndServe(":4000", r))
}

// Middleware, helper in the separate file

// isEmpty checks if the course is empty
func (c *Course) isEmpty() bool {
	return c.CourseName == ""
}

// Controllers (eventually stored in a separate file)

// serveHome serves the home page of the application
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome to the API by LearnCodeOnline</h1>"))
}

// getAllCourses gets all courses
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses ")                             // its like console.log statement 
	w.Header().Set("Content-Type", "application/json")          // we are setting the content type ie.json 
	json.NewEncoder(w).Encode(courses)                          // we have to encode our data in the slice to the json hence we have used json encoder we are 
	                                                            // writing the data to the frontend hence we have to pass the instance of http.responseWrite class
																// ie. "w"
}

// getOneCourse gets a course by ID
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r) // need to grab id from the request for finding the particular record in our fake database
	                      // all the parameters that will pass to the url will be stored in the params 

	// loop through the courses based on provided id and find the particular course 
	for _, course := range courses {            
		if course.CourseId == params["id"] {   // if we found desired id as provided by the user in url we will send the response to the user 
			                                   // via http.responseWriter class ("W") in the form of encoded json 
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course found with given id ")
	
}

// createOneCourse creates a new course
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course ")
	w.Header().Set("Content-Type", "application/json")
    
	// what if : body is empty
	if r.Body == nil {    //we are posting(sending ) some data to the specific url via the http.request method (r) 
		json.NewEncoder(w).Encode("Please send some data")  // thsi is the syntax for sending the information from backend to the frontend ie. (newEncoder(w).Encode)
		return
	}
    
	// what about - {}
	var course Course   // we have made the new slice of type Course 
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.isEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON ")
		return
	}
    
	// generate unique id which will be int covert it into the string and append to the courese slice 
	rand.Seed(time.Now().UnixNano())  // this syntax is used for creating the new id for inserting into the slice 
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)

	
}

// updateOneCourse updates a course
func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	 // first thing is grab id from the request 

	params := mux.Vars(r)


	// loop through the value if we find the id to be updated we will remove that recode and update the record served by the id 
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)

			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course) // will write the new details in the body i.e. (r.body) and will create a new element of type slice and 
			                                            // will append to the current slice 

			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("Course not found")
	
}

// deleteOneCourse deletes a course
func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
   
	
	// loop , id, remove(index, index+1)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}
}

// Controller handles requests on the particular route

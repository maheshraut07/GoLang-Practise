package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
)

// the Stud struct is defined with specific field tags using backticks () and the json` tag.
//These tags are used to provide metadata about how the struct fields should be serialized and deserialized when converting between Go structs and JSON.

type Stud struct {
	EnrollmentNumber string `json:"enrollment_number"`
	Name             string `json:"name"`
	Age              int    `json:"age"`
	Class            string `json:"class"`
	Subject          string `json:"subject"`
	IsDeleted        bool   `json:"is_deleted"`
}

var session *gocql.Session

func init() { // it is the function which will be called before the main function
	// Connect to Cassandra cluster
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "student"
	var err error

	session, err = cluster.CreateSession()
	if err != nil {
		log.Fatal("Error connecting to Cassandra: ", err)
	}
}

func main() {

	fmt.Println("Hello Friends welcome to the Application where we will be handling the students data using cassandra database ")

	r := mux.NewRouter()
	r.HandleFunc("/students", getAllStudents).Methods(http.MethodGet)
	r.HandleFunc("/students", addStudent).Methods(http.MethodPost)
	

	log.Fatal(http.ListenAndServe(":8080", r))
}

func getAllStudents(w http.ResponseWriter, r *http.Request) {

	var students []Stud

	query := "SELECT  enrollment_number, name, age, class, subject FROM stud WHERE is_deleted = false"

	iter := session.Query(query).Iter() //	Iter is used to iterate over the database
	var name, class, subject, enrollmentNumber string
	var age int

	//  the Scan method of the Iter object is used to
	//scan values from the database into variables name, age, classs, subject, and enrollmentNumber.
	// Since Scan expects pointers to variables, the & symbol is used to get the memory address of these variables.

	for iter.Scan(&enrollmentNumber, &name, &age, &class, &subject) {
		students = append(students, Stud{EnrollmentNumber: enrollmentNumber, Name: name, Age: age, Class: class, Subject: subject})
	}

	if err := iter.Close(); err != nil {
		http.Error(w, "Error retrieving students from Cassandra", http.StatusInternalServerError)
		return
	}

	//json.Marshal method is used to serialize a Go data structure into its JSON representation.
	//It takes a value of any type and returns a byte slice ([]byte) containing the JSON encoding of the value.
	jsonResponse, err := json.Marshal(students)

	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}

func addStudent(w http.ResponseWriter, r *http.Request){
	var stud Stud;

	err := json.NewDecoder(r.Body).Decode(&stud)

	if err!= nil{
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	// Generate UUID for enrollement Number

	stud.EnrollmentNumber = gocql.TimeUUID().String()

	query := `INSERT INTO stud (name, age , class, subject, enrollment_number,is_deleted) VALUES (?,?,?,?,?,false)`

	_ = session.Query(query,stud.Name,stud.Age,stud.Class,stud.Subject,stud.EnrollmentNumber).Exec()

	// return the UUID as response
}

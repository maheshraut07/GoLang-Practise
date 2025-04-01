package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Define a struct to parse the JSON request
type Course struct {
	CourseName string `json:"coursename"`
	Price      int    `json:"price"`
	Platform   string `json:"platform"`
}

func main() {
	fmt.Println("Starting the server on port 8000...")

	// Register handler for the /post route
	http.HandleFunc("/post", handlePostRequest)

	// Start the server
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

// Handler function for POST requests
func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close() // Close the body after reading

	// Parse the JSON data
	var course Course
	err = json.Unmarshal(body, &course)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Print received data on the server side
	fmt.Printf("Received course: %+v\n", course)

	// Respond with success message
	response := map[string]string{
		"message":   "Data received successfully",
		"coursename": course.CourseName,
		"price":      fmt.Sprintf("%d", course.Price),
		"platform":   course.Platform,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

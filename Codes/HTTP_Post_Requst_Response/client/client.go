package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Sending POST request to server...")

	const myurl = "http://localhost:8000/post"

	// Data to be sent as JSON
	requestBody, _ := json.Marshal(map[string]interface{}{
		"coursename": "Let's go with Golang",
		"price":      0,
		"platform":   "LearnCodeOnline.in",
	})

	// Making the POST request
	response, err := http.Post(myurl, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer response.Body.Close() // Close response body

	// Read the response
	var responseData map[string]interface{}
	json.NewDecoder(response.Body).Decode(&responseData)

	// Print the response
	fmt.Println("Response from server:", responseData)
}

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to the server")

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":5000", r)
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello, welcome to the backend of the GoLang server!</h1>"))
}

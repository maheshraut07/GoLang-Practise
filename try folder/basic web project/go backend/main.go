package main

import (
    "encoding/json"
    "log"
    "net/http"
)

type Message struct {
    Text string `json:"text"`
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
    message := Message{Text: "Hello from Go backend!"}
    json.NewEncoder(w).Encode(message)
}

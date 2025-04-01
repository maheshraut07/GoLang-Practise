// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func main() {
    
    // const myurl string = "https://jsonplaceholder.typicode.com/todos/1"
    const myurl string = "https://localhost/3000"
    
    response,err := http.Get(myurl)
    
    if err != nil{
        errors.New("some error is coming: ")
        log.Fatal(err)
                
    }
    
    fmt.Println("the response is : ", response)
    fmt.Println("the response body is : ", response.Body)
}
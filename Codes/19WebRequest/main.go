package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url string  = "https://jsonplaceholder.typicode.com/todos/1"

func main(){
	fmt.Println("it is the lecture for the web request")

	

	response , err := http.Get(url)  // get request of golang its like fetch api in the react 

	if err!= nil{
		panic(err)
	}

	fmt.Printf("the type of the response is : %T\n",response)


	// we can read the response by using ioUtil Package

	databytes, err := ioutil.ReadAll(response.Body)
	fmt.Println("this is the response in databyets form read using ioutil.ReadAll method: ", databytes)

	if err!= nil{
		panic(err)
	}

	content := string(databytes)
	fmt.Println("the content of the webpage is : \n", content)

	defer response.Body.Close()  // it is the user responsibilty to close 
}


/*

1) log.Fatal(err) prints the error and exits the program.

Why Use?
It provides an automatic timestamp with the error.
Exits the program immediately after an error.





2) Using errors.New() for Custom Errors

if err != nil {
        return errors.New("failed to fetch the URL, please check the server")
    }

Why Use?
Provides custom error messages instead of raw system errors.





3) Using log.Println(err) Instead of fmt.Println(err)
log.Println() prints errors with timestamps but does not stop execution.

if err != nil {
        log.Println("Warning: Unable to fetch URL:", err) // Logs error but continues execution
    } else {
        log.Println("Response received:", response.Status)
    }

Helps in debugging without crashing the entire application.


4)  Using recover() to Prevent Crashes
If an error panics the program, recover() prevents it from stopping completely.

unc main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from error:", r)
        }
    }()

    const myurl = "http://localhost:3000"

    response, err := http.Get(myurl)
    if err != nil {
        panic("Server is unreachable!") // Panic if an error occurs
    }
    
    fmt.Println("Response received:", response.Status)
}




Method	                     Description                            	Stops Execution?

fmt.Println(err)	  Simple error printing	                                ❌ No
log.Fatal(err)	      Prints error and exits	                            ✅ Yes
errors.New()	      Custom error message	                                ❌ No
log.Println(err)	  Logs error with timestamp	                            ❌ No
recover()	          Prevents panic crashes	                            ❌ No


*/
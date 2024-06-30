package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux" // for creating the server and handling the routes using router
)

func main() {
	fmt.Println("welcome to the modules (mod) in golang")

	greeter()

	// by using this syntax we handle the router (r)

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r)) // we have created the server that will run on the localhost and on the 4000 port number
	                                           // we bind it inside the log.fatal() method because if it has some error then it will handle all these erros 
}

func greeter() {
	fmt.Println("Hey there mod users")
}

// this is the standard method for handling the request and send the response

func serveHome(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h1> Welcome to the golang series </h1>")) // because we get the data in the form of bytes
	
}


/*


Gorilla Mux is a popular HTTP router for Go (Golang) web applications. It's a package within the Gorilla web toolkit, 
which provides various utilities for building web applications in Go. Mux specifically helps in routing incoming HTTP requests
to their respective handler functions based on the URL path and other request parameters.

Here's why and how Gorilla Mux is used:

Routing: Gorilla Mux helps in defining routes for different HTTP methods (GET, POST, PUT, DELETE, etc.) and URL patterns. This allows developers to create
clean and structured APIs.
URL Parameters: Mux allows capturing variables from the URL path, which can be used in the handler function to customize the response based on the request.
Middleware: It supports middleware, which are functions that can be executed before or after a request is handled. Middleware can be used for authentication,
logging, etc.
Subrouters: Mux supports subrouters, which can be used to create modular and reusable route configurations.
Customization: Gorilla Mux provides flexibility to customize routing behavior and handle edge cases effectively.
Here's a basic example of how Gorilla Mux is used:

go
Copy code
package main

import (
    "net/http"
    "github.com/gorilla/mux"
    "fmt"
)

func main() {
    // Create a new instance of the Mux router
    r := mux.NewRouter()

    // Define a route and its handler function
    r.HandleFunc("/hello/{name}", func(w http.ResponseWriter, r *http.Request) {
        // Get the value of the "name" variable from the URL path
        vars := mux.Vars(r)
        name := vars["name"]
        
        // Send a response with the greeting
        fmt.Fprintf(w, "Hello, %s!", name)
    })

    // Start the HTTP server with the Mux router
    http.ListenAndServe(":8080", r)
}
In this example, when a request is made to /hello/John, the handler function will extract the name "John" from the URL path and respond with "Hello, John!".



*/



// go mod tidy = will remove all the packages that are you not using 
                 

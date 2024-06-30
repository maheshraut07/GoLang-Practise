package main

import "fmt"

var Token int = 45 // variable name starting with the capital letter in means it it public 

func main(){
	var username string = "mahesh"
	fmt.Println(username)
	fmt.Printf("the type of the variable username is :  %T \n", username)
	fmt.Println()

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("the type of the variable username is :  %T \n", isLoggedIn)
	fmt.Println()

	var smallval uint8 = 255
	fmt.Println(smallval)
	fmt.Printf("the type of the variable username is :  %T \n", smallval)
	fmt.Println()

	var intval int = 255
	fmt.Println(intval)
	fmt.Printf("the type of the variable username is :  %T \n", intval)
	fmt.Println()

	var floatval float32 = 255.94837320952  // 5 values after the decimal in case of float 64 more than 5
	fmt.Println(floatval)
	fmt.Printf("the type of the variable username is :  %T \n", floatval)
	fmt.Println()

	//default values and some alises 

	var anotherVar int 
	fmt.Println(anotherVar)
	fmt.Printf("the type of the variable username is :  %T \n", anotherVar)
	fmt.Println()

	// implicit type

	var website = "mahi.com"
	fmt.Println(website)
	fmt.Println()

	// no var style

	numberOfUser := 30000.0
	fmt.Println(numberOfUser)
	fmt.Println()

	// accessing the public variable 
	fmt.Println(Token)







	

}

/*
*** uint8 = unsigned 8-bit integer (0 t0 255)
uint16
uint32
unit64

int8 = signed 8-bit integer (-128 to 127)




*****In Go, the := operator is used for variable declaration and assignment. It is a shorthand notation for declaring and initializing variables. Here's how it works:

When you use :=, Go infers the type of the variable from the assigned value.
It can only be used inside functions, not at the package level.
It automatically declares a new variable if it doesn't exist, or reuses an existing variable if it does.
Here's an example:

go
Copy code
package main

import "fmt"

func main() {
    // Using :=
    name := "Alice"
    age := 30
    fmt.Println(name, age)

    // Using =
    var city string
    city = "New York"
    fmt.Println(city)
}
In the above example, name and age are declared and initialized using :=, while city is declared using =. The = operator is used for assignment without type inference. It can be used both inside and outside functions, and it assigns a value to a variable.


*/
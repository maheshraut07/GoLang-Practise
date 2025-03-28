package main

import "fmt"

// here how  you can  declare the structure

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("Structs in golang")

	// here how you can access the struct 

	mahesh := User{"mahesh","rautmahesh213@gmail.com",true,21}
	fmt.Println(mahesh)

	fmt.Printf("mahesh details are : %+v\n", mahesh)
	fmt.Println()
	fmt.Printf("name is %v and email is %v", mahesh.Name, mahesh.Email)
	fmt.Println()

}

// there is no inheritance in golang as there are no classes

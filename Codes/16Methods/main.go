package main

import "fmt"

// Declaration of the User struct
type User struct {
    Name   string
    Email  string
    Status bool
    Age    int
}

// Method to get the status of the user
func (u *User) GetStatus() {
    fmt.Println("the status is : ", u.Status)
}

// Method to update the email of the user
func (u *User) NewEmail() {
    u.Email = "maheshraut8528@gmail.com"
    fmt.Println("Email of this user is : ", u.Email)
}

func main() {
    fmt.Println("Structs in golang")

    // Creating an instance of the User struct
    mahesh := User{"mahesh", "rautmahesh213@gmail.com", true, 21}
    fmt.Println(mahesh)

    // Accessing struct fields
    fmt.Printf("mahesh details are : %+v\n", mahesh)
    fmt.Println()
    fmt.Printf("name is %v and email is %v", mahesh.Name, mahesh.Email)
    fmt.Println()

    // Calling the GetStatus method
    mahesh.GetStatus()
    fmt.Println()

    // Calling the NewEmail method
    mahesh.NewEmail()
    fmt.Println()
}

/*

Methods in Golang are functions with a receiver. They allow you to attach behaviors (functions) to a struct.

Method 1: GetStatus()
go
Copy
Edit
func (u *User) GetStatus() {
    fmt.Println("the status is : ", u.Status)
}
func (u *User) GetStatus() defines a method associated with the User struct.

The (u *User) is called a receiver, which means this function operates on a User struct.

This method prints the Status of the user.

Method 2: NewEmail()
go
Copy
Edit
func (u *User) NewEmail() {
    u.Email = "maheshraut8528@gmail.com"
    fmt.Println("Email of this user is : ", u.Email)
}
This method modifies the Email field of the struct.

Since u is a pointer (*User), changes to u.Email persist outside the method.

3️⃣ Using Struct in main()
go
Copy
Edit
mahesh := User{"mahesh", "rautmahesh213@gmail.com", true, 21}
Creates a new User instance named mahesh and initializes it with values.

This is struct literal notation.

go
Copy
Edit
fmt.Printf("mahesh details are : %+v\n", mahesh)
%+v prints the struct with field names.

go
Copy
Edit
fmt.Printf("name is %v and email is %v", mahesh.Name, mahesh.Email)
Accessing struct fields using dot notation (mahesh.Name, mahesh.Email).

go
Copy
Edit
mahesh.GetStatus()
mahesh.NewEmail()
Calls the GetStatus() and NewEmail() methods.

Struct Syntax in Golang
1️⃣ Basic Struct Declaration
go
Copy
Edit
type StructName struct {
    Field1 Type
    Field2 Type
}
2️⃣ Creating a Struct
go
Copy
Edit
var user1 User = User{"Alice", "alice@gmail.com", true, 25}
OR

go
Copy
Edit
user1 := User{Name: "Alice", Email: "alice@gmail.com", Status: true, Age: 25}
3️⃣ Accessing Struct Fields
go
Copy
Edit
fmt.Println(user1.Name)
4️⃣ Defining Methods
go
Copy
Edit
func (u *User) PrintName() {
    fmt.Println("User name is:", u.Name)
}
5️⃣ Calling Methods
go
Copy
Edit
user1.PrintName()
Methods of Struct in Golang
Method	Description
func (s StructName) MethodName()	Defines a struct method with a value receiver.
func (s *StructName) MethodName()	Defines a struct method with a pointer receiver (allows modification of struct fields).
fmt.Printf("%+v", structName)	Prints struct fields and values.
structName.FieldName	Access a specific field of the struct.
Conclusion
A struct in Go is a collection of related data stored as fields.

Methods allow behavior to be attached to structs.

Use pointer receivers (*User) when modifying struct fields.

Structs are widely used for modeling real-world data in Golang.

*/

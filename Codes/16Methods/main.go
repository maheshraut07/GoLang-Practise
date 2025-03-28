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

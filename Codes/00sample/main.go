package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    
    
    var subjects = []string{"maths", "english", "science"}
    var names = [3]string{"mahesh", "sumit", "vicky"}

    fmt.Printf("The type of the names array is: %T\n", names)
    fmt.Println()
    fmt.Printf("The type of the subject slice is: %T\n", subjects)
    fmt.Println()
    fmt.Println("This is slice:", subjects)

    fmt.Println("The length of the subjects is:", len(subjects))
    fmt.Println("The first two elements in the subjects are:", subjects[0:2])
    fmt.Println()
    
    
	reader := bufio.NewReader(os.Stdin)
    fmt.Println("enter the any number: ")
    input, _ := reader.ReadString('\n')
    fmt.Println("the inputted number is : ", input)
}

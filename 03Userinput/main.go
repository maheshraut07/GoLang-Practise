package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){

	welcome := "welcome to user Input"
	fmt.Println(welcome)

	// this is the syntax for reading user input and output 

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza : ")

	// in go lang there are no try and catch block 
	// so you are accepting the input from the user in a variable and if some error occurs you accept in in the "_" or "err"
    // so in this case we use "comma ok" or "comma error" syntax 

	
	input, _ := reader.ReadString('\n')  //When ReadString encounters the newline character \n in the input, it stops reading and returns the string read up to that point, including the newline character itself.
	fmt.Println("Thanks for rating , the rating is : " , input )
	fmt.Printf("the type of the input variable is: %T",input)
}
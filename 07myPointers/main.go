package main 

import "fmt"

func main (){
	fmt.Println("Welcome to a class on the pointers")
	
	var ptr *int 
	fmt.Println("The value of the pointer is : ", ptr)
	fmt.Println()

	myNumber := 23
	var ptr1 = &myNumber
    fmt.Println("the address of the variable muNumber is : ",ptr1)
    fmt.Println("the address of the variable muNumber is : ",*ptr1)
	fmt.Println()

	*ptr1 = *ptr1 + 2
	fmt.Println("The modified value of the myNumber is : ", *ptr1)

}
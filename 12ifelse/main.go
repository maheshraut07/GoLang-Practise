package main

import "fmt"

func main(){
	fmt.Println("This is the lecture for if else ")


	logicount := 23
	var result string 

	if logicount < 10{
		result = "Regular User"
	}else if logicount > 10{
		result = "watch out "
	}else{
		result = "exactly 10 times "
	}

	fmt.Println("teh value of the user is : ", result)
	fmt.Println()


	// another way 

	if 9%2 == 0{
		fmt.Println("the number is even")
	}else{
		fmt.Println("the number is odd")
	}

	fmt.Println()

	// another way

	if num:= 3; num < 10 {
		fmt.Println("The num is less than 10")
	}else{
		fmt.Println("the num is greater than 10")
	}
}


// we can avoid the use of the parenthesis in if conditions 
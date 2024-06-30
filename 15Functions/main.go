package main

import "fmt"

func main(){

	
	fmt.Println("welcome to functions in the golang")

	// you cant define the function inside the another function 

	greeter()
	greetertwo(2,3)
	fmt.Println()

	result := adder(2,3)
	fmt.Println("the addtion is : ", result)
	fmt.Println()

	fmt.Println("the addtion of first 10 natural numbers are : ",proAdder(1,2,3,4,5,6,7,8,9,10))
	fmt.Println()

	// for handling the proAdder2 function
	proResult , myMeassage := proAdder2(1,2,3,4,5,6,7,8,9,10)
	fmt.Println("the addtion by using proAdder is : ", proResult)
	fmt.Println("the message while doing proAdder is : ",myMeassage)
	fmt.Println()


}


func greeter (){
	fmt.Println("hii from the greeter function")
}

func greetertwo( a int ,  b int){

	fmt.Println("hii from the greetertwo function")
	fmt.Printf("the addition of %v and %v is %v  ",a,b,a+b)
}

func adder(a int, b int) int {  // this is the return type of the function 
	return a + b
}

// if you dont know how many parameters are inside the function then you can define the following function

func proAdder(values ...int)int {
     
	total := 0;

	for _ , val := range values{
		total += val
	}

	return total
}

// there is another method for returning multiple values 

func proAdder2(values ...int)(int, string) {
     
	total := 0;

	for _ , val := range values{
		total += val
	}

	return total , "hii the addition is calculated "
}


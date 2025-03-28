package main

import "fmt"

func main(){
	fmt.Println("this lecture is for the loops")

	days := []string {"sunday","monday","tuesday"}
	fmt.Println("The days slice is : ", days)

	// this is the syntax for the loops 

	for d:= 0; d < len(days); d++{
		fmt.Println("the days are : ",days[d])
	}
	fmt.Println()

	// there is another way for performing the opeartions on the loops 

	for i := range days{
		fmt.Println("the days are :",days[i])
	}
	fmt.Println()

	// there is for each type of loop in golang and the syntax for the same is given below 

	for index,value := range days{
		fmt.Printf("index is %v and the value is %v", index,value)
		fmt.Println()
	}
	fmt.Println()

	// there are no loops like while loops in the golang but you can do the same by using the following syntaxx 

	tempValue := 1;

	for tempValue < 10{
       
		if tempValue == 2{
			goto mahi
		}

		if tempValue == 5 {
			tempValue++
			continue
		}

		if tempValue == 8{
			break
		}

		fmt.Println("the value is : ",tempValue)
		tempValue++
	}
	fmt.Println()

	mahi : 
		fmt.Println("this is the mahi label ")

}
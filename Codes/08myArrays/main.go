package main

import (
	"fmt"
)

func main (){
	fmt.Println("welcome to the lecture of the myarray ")
	fmt.Println()

    var fruitlist [4]string

	fruitlist[0] = "Apple"
	fruitlist[1] = "Tomato"
	fruitlist[3] = "Peach"

	fmt.Println("Fruit list is : ", fruitlist)  // there would be the space between first and third item as it is not mentioned 
	fmt.Println("the length of the Fruit list is : ", len(fruitlist))
	fmt.Println()

   
	var vegList = [3] string{"potato","beans","mushroom"}
    fmt.Println("veglist is : ",vegList)
	fmt.Println()



	


}
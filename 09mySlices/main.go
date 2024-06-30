package main

import (
	"fmt"
	"sort"
)

func main(){
	fmt.Println("welcome to the lecture on the myslices")

	var fruitlist = [] string {"Apple", "Banana", "WaterMelon"}
	fmt.Printf("the type of the fruitlist is : %T", fruitlist) // it will print the []string that is the identity of the slice operator 
	fmt.Println()

	fruitlist = append(fruitlist, "Peach","Mango")
	fmt.Println("the modified fruitlist is : ",fruitlist)
	fmt.Println()

	fmt.Println("the fruitlist after performing slice opearation is : ", fruitlist[1:3])
	fmt.Println()

	//fruitlist = append(fruitlist[2:])  // it will modify the fruitlist slices 
	fmt.Println("the modified fruitlist is : ", fruitlist)



	// use of make keyword 

	highScores := make([]int, 4)   // it will the make slice of 4 integer 
	
	highScores[0] = 10;

	fmt.Println("the highscores array is : ", highScores)

	highScores = append(highScores, 6,5,8)
	fmt.Println("The modified highscores array is : ", highScores)
	fmt.Println()

	// sort method 

	sort.Ints(highScores)
	fmt.Println("the highscores array after sorting is : ", highScores)
	fmt.Println()

	fmt.Println("the highscores array is sorted or not : ", sort.IntsAreSorted(highScores))
	fmt.Println()


	// how to remove the value from slices based on index 

	var courses = [] string{"reactJs", "javascript","swift","ruby","python"}
	fmt.Println("the courses array is : ",courses)
	fmt.Println()

	var index int = 2 
	//wrong syntax 
	//courses = append(courses[:index], courses[index+1:])
	courses = append(courses[:index],courses[index+1:]...)
	fmt.Println("the courese array after modification is : ",courses)


}









// slices are the advancements of the array 
// slices are just like the dynamic arrays 
// we can avoid to give the size to the slices in go lang 
// we can add and remove values in the slices by using append method 
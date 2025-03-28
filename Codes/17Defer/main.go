package main

import "fmt"

func main(){
	defer fmt.Println("World")  // here we used the defer keyword here so this line would execute as the last line of the  function 
	fmt.Println("Hello")
	fmt.Println()

	DeferUse()  // it will be printed in last in first out 

}

func DeferUse(){
	defer fmt.Println("one")
	defer fmt.Println("Two")
	defer fmt.Println("Three")
}

// differ opearates in LIFO mode 
// it will create mcdonlad queue 

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	fmt.Println("this is the lecture for switch case ")

	// we will create a random number 

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
    fmt.Println("value of the dice is : ",diceNumber)

	switch (diceNumber){
	case 1 :
		fmt.Println("Dice value is 1 and you can open")

	case 2 :
		fmt.Println("you can move to 2 spot ")
		fallthrough    // it is kind of the continue s
	case 3 :
		fmt.Println("you can move to 3 spot ")
	case 4 :
		fmt.Println("you can move to 4 spot ")
	case 5 :
		fmt.Println("you can move to 5 spot ")
	case 6 :
		fmt.Println("you can move to 6 spot and roll dice again  ")
	default:
		fmt.Println("what was that!! ")
	}
}
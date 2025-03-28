package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pizza in the rating between 0-9 :")

	reader := bufio.NewReader(os.Stdin)

	input , _ := reader.ReadString('\n')
	fmt.Println("thanks for rating the pizza the rating is : " , input)
	fmt.Printf("the type of the input variable is : %T", input)
	fmt.Println()


	// the input type is string so you have to covert it for doing the mathematical opeartion 

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input),64) // it will take the two parameters the input string and the number of bits you want to convert into 

	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(numRating)
		fmt.Println("Added 1 to your rating : ", numRating + 1)
		fmt.Printf("the type of the converted input is : %T",numRating)
	}

}
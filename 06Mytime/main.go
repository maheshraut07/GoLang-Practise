package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("Welcome to time study of golang")

    // the following syntax for fetching the current date 
	presentTime := time.Now()

	fmt.Println("The present time is : ", presentTime)
	fmt.Println()

	// if you want to format it the you can use format function of time 

	fmt.Println("the present time is :", presentTime.Format("01-02-2006  15:04:05 Monday"))

	// the following syntax for creating the date 

	createdDate := time.Date(2003,time.August,9,12,23,0,0,time.UTC)
	fmt.Println("the created date is :", createdDate.Format("01-02-2006  15:04:05 satuarday"))
}
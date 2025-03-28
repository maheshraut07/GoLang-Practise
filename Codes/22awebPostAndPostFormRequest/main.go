package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main(){

	fmt.Println("Welcome to the lecture of the post requests ")

	performPostRequest()

}

func performPostRequest(){
	const myurl = "http://localhost:8000/post"

	requestBody := strings.NewReader(`
	{
		"coursename":"Let's go with golang",
		"price":0,
		"platform":"LearnCodeOnline.in"
	}
	
	`)

	// so we are going to make the post request to our url 

	response , err := http.Post(myurl, "application/json",requestBody)

	if err != nil{
		panic(err)
	}

	// for reading the data sent by the post request we need the ioUtil package 



	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println("the content is : ", string(content ))
    

	defer response.Body.Close()

}



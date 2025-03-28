package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main(){
	fmt.Println("welcome to the handling of web requests")
	PerformGetRequest()
}

func PerformGetRequest(){
	const myurl = "http://localhost:3000/"

	response , err := http.Get(myurl)

	if err!= nil{
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("the status code is : ", response.StatusCode)
	fmt.Println("Content length is : ", response.ContentLength)

	// i want to read the all content specified by the url 

	// there is standard strings library for reading the responses

	var responseString strings.Builder

	content , _ := ioutil.ReadAll(response.Body)   // the content will be in the byte format 
	bytecount, _ := responseString.Write(content)
	
	fmt.Println("Bytecount is : ",bytecount)
	fmt.Println("the actual data in the response is : ",responseString.String())


	// so we will convert it first in the format of the string 
	

	fmt.Println("the actual data in the response is : ", string(content))

}



// in golang string is the datatype and strings is the package 

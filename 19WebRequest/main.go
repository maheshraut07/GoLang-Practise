package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url string  = "https://courses.learncodeonline.in/"

func main(){
	fmt.Println("it is the lecture for the web request")

	

	response , err := http.Get(url)  // get request of golang its like fetch api in the react 

	if err!= nil{
		panic(err)
	}

	fmt.Printf("the type of the response is : %T\n",response)


	// we can read the response by using ioUtil Package

	databytes, err := ioutil.ReadAll(response.Body)

	if err!= nil{
		panic(err)
	}

	content := string(databytes)
	fmt.Println("the content of the webpage is : \n", content)

	defer response.Body.Close()  // it is the user responsibilty to close 
}
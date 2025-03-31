package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://loc.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"
// const myurl string = "https://jsonplaceholder.typicode.com/todos/1"

func main (){
	fmt.Println("Welcome to the lecture of handling URLs")
	fmt.Println(myurl)

	// for parsing the url there  is another package called as url 

	result , _ := url.Parse(myurl)
	fmt.Println(result.Scheme)  
	fmt.Println(result.Host)  
	fmt.Println(result.Path)  
	fmt.Println(result.RawQuery)  
	fmt.Println(result.Port())  

	// store the params 

	qparams := result.Query()
	fmt.Printf("the result of the query params are : %T\n ",qparams)

	fmt.Println(qparams["coursename"])
	fmt.Println(qparams["paymentid"])
	fmt.Println()

	for _, value := range qparams{
		fmt.Println("param is : ", value)
	}
	fmt.Println()

    // constructing the URLs 
	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println("another url is : ",anotherURL)


}
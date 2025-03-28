package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main(){
	fmt.Println("Welcome to the file ")

	// there is os module to create the files 

	content := "This needs to go in a file "

	file , err := os.Create("./myfile.txt")

	checkNilErr(err)

	

	// for writing into the file we have io package 

	length, err := io.WriteString(file , content)

	checkNilErr(err)

	fmt.Println("The length is : ",length)
	file.Close()

	readFile("./myfile.txt")

}

func readFile(filename string ){
	databyte , err := ioutil.ReadFile(filename) // for reading the file we have ioUtil package

	checkNilErr(err)

	content := string(databyte)

	fmt.Println("Text data inside the file is :  \n", content)
}

func checkNilErr (err error){
	if err!= nil{
		panic(err)  // panic will shut down the execution of the program 
    }
}

// for creating the  files there is io utility
// for reading those files  we have ioUtils utility

// files are being read in a databyte format not in a string format 
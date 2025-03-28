package main

import (
	"encoding/json"
	"fmt"
)

// we are going to construct the structure

type course struct{
	Name string       `json:"courseName"`          // we can change the name of the name in the json format by using this syntax 
	Price int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags []string     `json:"tags,omitempty"`  // this is the slice 
	                                           // if it is empty it is not going to show 
}


func main(){
	fmt.Println("Hi this is the lecture for Json data ")

	// call the function EncodeJson inside the main function 

	fmt.Println("the starting of EncodeJson function")
	EncodeJson()
	fmt.Println()
	fmt.Println("the starting of DecodeJson function")
	fmt.Println()
	DecodeJson()
} 

func EncodeJson(){
	// this function will have the content based on the struct type 

	lcoCourses := [] course{
		{"ReactJS BootCamp",299,"LearnCodeOnline.in","abc123",[]string{"web-dev","js"}},
		{"MERN BootCamp",399,"LearnCodeOnline.in","abc456",[]string{"full-stack","js"}},
		{"DSA BootCamp",499,"LearnCodeOnline.in","abc789",nil},
		
	}

	finalJson, err := json.MarshalIndent(lcoCourses,"","\t")  //data is indented based on the passed parameter 

	if err != nil{
		panic(err)
	}

	fmt.Printf("%s\n",finalJson)


}

func DecodeJson(){

	jsonDataFromWeb := []byte(`
	    {
			"courseName": "ReactJS BootCamp",
			"Price": 299,
			"website": "LearnCodeOnline.in",
			"tags": ["web-dev","js"]
	    }
		`)

	var lcoCourses course

	isvalidJson := json.Valid(jsonDataFromWeb)  // to check if the data is valid or not 

	if isvalidJson{
		fmt.Println("the json data is vaild")
		json.Unmarshal(jsonDataFromWeb, &lcoCourses) // for not making the copy of the data 
		fmt.Printf("%#v\n",lcoCourses)   // it will print the json data in the format of the struct 
    }else{
		fmt.Println("the json was not valid ")
	}


	// in some cases you just wanted to add the data to key value 
	fmt.Println()

	var myonlineKeyValueData map[string]interface{}  // the key will be in the string format but we dont know what the foramt of the value thats why the interface type is 

	json.Unmarshal(jsonDataFromWeb, &myonlineKeyValueData)
	fmt.Printf("%#v\n",myonlineKeyValueData)

}



// the data that comes from the web is in the byte format so we have 
// convert it in the string of json format for rading it 
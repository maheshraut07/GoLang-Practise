package main

import (
	"fmt"
)

func main(){
	fmt.Println("This is the lecture for Maps")

	languages := make (map[string]string,7)

	languages["JS"] = "javascript" 
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of the all languages is : ", languages)
	fmt.Println("shortcut for js is :",languages["JS"])
    

	delete(languages,"RB") // this is how you can delete the values 


	// for loops in case of maps 

	for key , value := range languages{
		fmt.Printf("For key %v , value is %v\n", key, value)

	}


}

    
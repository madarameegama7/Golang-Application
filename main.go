package main

import "fmt"

func main(){
	fmt.Println("Hello, Madara")

	//strings
	//if we declare a variable it must be used otherwise we get an error
	var nameOne string = "mario"
	var nameTwo = "luigi"
	var nameThree string


	fmt.Println(nameOne)
	fmt.Println(nameTwo,nameThree)

	nameOne = "dinu"
	nameThree = "vidura"

	fmt.Println(nameOne,nameTwo,nameThree)

	//short method to declare variable(we cannot use this outside function)

	nameFour := "nayana"
	fmt.Println(nameFour)


}


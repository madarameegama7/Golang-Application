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

	//ints
	var ageOne int =25
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne,ageTwo,ageThree)

	//bits and memory
	var numOne int8 = 125
	var numTwo int8 = -128
	var numThree uint8 = 25

	fmt.Println(numOne, numTwo, numThree)

	//float numbers
	var scoreOne float32 = -1.5
	var scoreTwo float64 = 8845157846.3
    scoreThree := 1.5

	fmt.Println(scoreOne, scoreTwo, scoreThree)



}


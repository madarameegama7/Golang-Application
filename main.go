package main

//this fmt package is used to format strings and print in console
import "fmt"

func main(){
	// fmt.Println("Hello, Madara")

	// //strings
	// //if we declare a variable it must be used otherwise we get an error
	// var nameOne string = "mario"
	// var nameTwo = "luigi"
	// var nameThree string


	// fmt.Println(nameOne)
	// fmt.Println(nameTwo,nameThree)

	// nameOne = "dinu"
	// nameThree = "vidura"

	// fmt.Println(nameOne,nameTwo,nameThree)

	// //short method to declare variable(we cannot use this outside function)

	// nameFour := "nayana"
	// fmt.Println(nameFour)

	// //ints
	// var ageOne int =25
	// var ageTwo = 30
	// ageThree := 40

	// fmt.Println(ageOne,ageTwo,ageThree)

	// //bits and memory
	// var numOne int8 = 125
	// var numTwo int8 = -128
	// var numThree uint8 = 25

	// fmt.Println(numOne, numTwo, numThree)

	// //float numbers
	// var scoreOne float32 = -1.5
	// var scoreTwo float64 = 8845157846.3
    // scoreThree := 1.5

	// fmt.Println(scoreOne, scoreTwo, scoreThree)

	// //Print
	// fmt.Print("hello, ")
    // fmt.Print("world! \n")
	// fmt.Print("world! \n")

	age := 22
	name := "madara"
	// fmt.Println("My age is", age, "and my name is", name)

	//Printf-formatted strings %_ = format specifier
	fmt.Printf("my age is %v and my name is %v \n", age, name)
	fmt.Printf("my age is %q and my name is %q \n", age, name)
	fmt.Printf("age is of type %T", age)
	fmt.Printf("you scored %0.1f points! \n", 225.55)
	
	//sprintf - saved formatted strings
	var str =fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("The saved string is: ",str)

	//arrays and slices
	//var ages [3] int = [3]int{1,2,3}
    var ages = [3]int{1,2,3}

	names := [3]string{"madara", "nayana", "sai"}

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))


}


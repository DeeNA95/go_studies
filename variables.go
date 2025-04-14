package main

import "fmt"

var DECORATOR = "=====================" // this is a global variable

func main(){
	// strings nb one method of declaring vars in go is "var var_name type = value"
	var firstName string = "Derek"

	// another method of declaring vars in go is "var var_name = value"
	var lastName  = "Amuna"

	// initialising without a value
	var age int
	// for ints and floats though it defaults to 0
	age = 20

	var gpa float32
	gpa = 3.5

	// also has memory allocation ie int8, float64, uint8, uint16, uint32, uint64

	var unsignedInt uint = 10 // unsigned int eg is positive 
	fmt.Println(unsignedInt)
	// another way of declaring vars

	isStudent := false //using the walrus operator but can only be done withing functions

	fmt.Println(DECORATOR)
	fmt.Println("Hello, my name is",firstName, lastName, "and I am", age, "years old.", "and my gpa is", gpa)
	fmt.Println("Am I a student?", isStudent)

}

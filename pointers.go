package main

import "fmt"

func updatename(name *string) {
	*name = "Bob"
}

func main() {
	name := "Alice"

	fmt.Println("memory address of name:", &name) // prints the memory address of the variable name

	// m := &name // m is a pointer to the variable name
	// fmt.Println("memory address of m:", m) // prints the memory address of the variable name
	// fmt.Println("value of m:", *m) // prints the value of the variable name

	// *m = "Bob" // dereference m and assign a new value to the variable name
	// fmt.Println("updating the value of m to Bob")
	// fmt.Println("value of name:", name) // prints the value of the variable name
	// fmt.Println("value of m:", *m) // prints the value of the variable name

	// now using the updatename btu with a pointer
	fmt.Println("Before:", name)
	updatename(&name) // pass the address of name to the function
	fmt.Println("After:", name) // name is now Bob
}

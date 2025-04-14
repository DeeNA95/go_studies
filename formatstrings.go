package main
import "fmt"
func main() {
	age:=22
	name:="Derek"

	// Print a string
	fmt.Print("This is a formatted string.") // Print without a newline
	fmt.Print("This is a formatted string.\n") // Print with a newline

	// Print a string with a newline
	fmt.Println("This is another formatted string.") // Print with a newlin
	fmt.Println("Hello, my name is", name, "and I am", age, "years old.")

	//Printf
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", name, age)
	fmt.Printf("Hello, my name is %v and I am %v years old.\n", name, age) //%v is a generic placeholder
	fmt.Printf("Hello, my name is %q and I am %d years old.\n", name, age) //%q is a quoted string
	fmt.Printf("Hello, my name is %T and I am %T years old.\n", name, age) //%T is the type of the variable
	fmt.Printf("this is a string with a float %f\n", 3.14) //%f is a float
	fmt.Printf("this is a string with a float %.2f\n", 3.14) //%0.2f is a float with 2 decimal places
	fmt.Printf("this is a string with a float %.2f\n", 3.14159265358979323846) //%0.2f is a float with 2 decimal places

	//Sprintf
	// Sprintf is used to format a string and return it as a string
	// it does not print the string to the console
	str := fmt.Sprintf("Hello, my name is %s and I am %d years old.", name, age)
	fmt.Println("==========", "This is the Sprintf version", "==========")
	fmt.Println(str) // Print the formatted string
}

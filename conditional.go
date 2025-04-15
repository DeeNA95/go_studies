package main

import "fmt"

func main(){
	age := 22
	fmt.Println(age > 18) // prints true
	fmt.Println(age < 18) // prints false
	fmt.Println(age == 18) // prints false
	fmt.Println(age != 18) // prints true
	fmt.Println(age >= 18) // prints true
	fmt.Println(age <= 18) // prints false
	fmt.Println(age > 18 && age < 30) // prints true && means and
	fmt.Println(age > 18 || age < 20) // prints true || means or
	fmt.Println(!(age > 18)) // prints false
	fmt.Println(age > 18 && age < 30 && age != 25) // prints true goes left to right

	if age > 18 {
		fmt.Println("You are an adult")
	}
	if age := 17; age < 18 { //scoped variable
		// this variable is scoped to the if statement
		fmt.Println("You are a minor")
	} else {
		fmt.Println("You are an adult")
	}
	fmt.Println(age)

	if age < 18{
		fmt.Println("You are a minor")
	} else if age > 18 && age < 30 {
		fmt.Println("You are a young adult")
	} else {
		fmt.Println("You are an adult")
	}
	names := []string{"Derek", "Amuna", "John", "Doe"}

	for _,name :=range names{
		if len(name ) <= 4{
			fmt.Println(name)
		} else {
			continue
		}
	}


}

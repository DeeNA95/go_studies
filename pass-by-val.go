package main

import "fmt"

func updatename(name string) {
	name = "Bob"
}

func updatemenu(mapping map[string]int) {
	mapping["one"] = 100
}

func main() {
	name := "Alice"
	fmt.Println("Before:", name)
	updatename(name)
	fmt.Println("After:", name) // name is still Alice
	// This is because strings are passed by value in Go, so the original string is not modified.
	// If you want to modify the original string, you need to pass a pointer to it. or do an assignments to the same var name

	// /some types dont adhere to this though
	// maps, slices, functions are passed by reference
	// so the original map is modified
	menu := map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
	}
	fmt.Println("Before:", menu)
	updatemenu(menu)
	fmt.Println("After:", menu) // menu is now updated
}

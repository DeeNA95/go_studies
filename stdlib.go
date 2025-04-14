package main

import (
	"fmt"
	"sort"
)

func main() {
	// greeting := "Hello, World!"

	// fmt.Println(strings.Repeat("=", 20)) //it auto imports the strings package
	// fmt.Println(strings.Contains(greeting, "World")) // prints true

	// fmt.Println(strings.ReplaceAll(greeting, "World", "Golang")) // prints Hello, Golang!
	// fmt.Println(strings.ToUpper(greeting)) // prints HELLO, WORLD!
	// fmt.Println(strings.ToLower(greeting)) // prints hello, world!
	// fmt.Println(strings.TrimSpace(greeting)) // prints Hello, World!
	// fmt.Println(strings.Trim(greeting, "!")) // prints Hello, World
	// fmt.Println(strings.TrimLeft(greeting, "H")) // prints ello, World!
	// fmt.Println(strings.TrimRight(greeting, "!")) // prints Hello, World
	// fmt.Println(strings.TrimPrefix(greeting, "Hello")) // prints , World!
	// fmt.Println(strings.TrimSuffix(greeting, "!")) // prints Hello, World
	// fmt.Println(strings.Split(greeting, ",")) // prints [Hello  World!]
	// fmt.Println(strings.Join([]string{"Hello", "World"}, ",")) // prints Hello,World
	// fmt.Println(strings.Index(greeting, "ll")) // prints 7

	ages := []int{30,55,20, 21, 22, 23, 24}

	sort.Ints(ages) // sorts the slice in ascending order
	fmt.Println(ages) // prints [20 21 22 23 24]

	index := sort.SearchInts(ages, 22) // searches for the index of 22 in the slice

	fmt.Println(index) // prints 2

	names := []string{"Derek", "Amuna", "John", "Doe"}
	sort.Strings(names) // sorts the slice in ascending order
	fmt.Println(names) // prints [Amuna Derek Doe John]
}

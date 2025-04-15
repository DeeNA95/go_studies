package main

import "fmt"

func main(){
	//maps: made up of key-value pairs
	// similar to python dicts
	//
	// has the format map[key]value where key and value are the types of the key and value

	// a map of string to int

	m := map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
		"four": 4,
	}
	fmt.Println(m)
	fmt.Println(m["one"]) // prints the map the value of the key "one"
	fmt.Println(m["five"]) // prints 0 because the key "five" does not exist in the map

	// looping maps
	for key, value := range m {
		fmt.Println(key,"–", value)
	}

	//keys and values can be any type
}

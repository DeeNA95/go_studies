package main
import "fmt"

func main() {
	//arrays have a fixed size
	var array1 [5]int // declaring an array of size 5
	var array2 [5]int = [5]int{1, 2, 3, 4, 5} // declaring and initializing an array
	var array3 = [...]int{1, 2, 3, 4, 5} // declaring and initializing an array with size inferred
	var array4 = [5]int{1: 10, 3: 30} // declaring and initializing an array with specific indices
	fmt.Println("Array 1:", array1, "has length", len(array1)) // prints [0 0 0 0 0]
	fmt.Printf("Array 1: %v has length %d\n", array1, len(array1)) // prints [0 0 0 0 0]
	fmt.Println("Array 2:", array2) // prints [1 2 3 4 5]
	fmt.Println("Array 3:", array3) // prints [1 2 3 4 5]
	fmt.Println("Array 4:", array4) // prints [0 10 0 30 0]
	fmt.Println("array4[1] prints the second value in the array ie", array4[1]) // prints 10
	//usual indexing and inputting
	array4[1] = 20 // changing the value at index 1
	fmt.Println("Array 4 after changing the value at index 1:", array4, "also done for slices") // prints [0 20 0 30 0]

	//slices are dynamic arrays
	var slice1 = []int{1,2,3} // declaring a slice
	var slice2 = []int{1, 2, 3, 4, 5} // declaring and initializing a slice
	fmt.Println("Slice 1:", slice1) // prints []
	fmt.Println("Slice 2:", slice2) // prints [1 2 3 4 5]
	// appending
	slice1 = append(slice1, 1) // appending to a slice
	fmt.Println("Slice 1 after appending:", slice1)

	//ranges
	range1 := slice2[1:3] // slicing a slice
	range2 := slice2[1:] // slicing a slice from index 1 to the end
	range3 := slice2[:3] // slicing a slice from the beginning to index 3
	range4 := slice2[:] // slicing a slice from the beginning to the end
	fmt.Println("Range 2:", range2) // prints [2 3 4 5]
	fmt.Println("Range 3:", range3) // prints [1 2 3]
	fmt.Println("Range 4:", range4) // prints [1 2 3 4 5]
	fmt.Println("Range 1:", range1) // prints [2 3]

}

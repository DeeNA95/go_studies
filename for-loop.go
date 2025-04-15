package main

func main(){
	// practically the same as PYTHON and CPP
	x := 0
	for x < 10 {
		println(x)
		x++
	}

	// // now for the cpp type
	for i:=0; i < 10; i++ {
		println(i)
	}

	names := []string{"Derek", "Amuna", "Derek Amuna"}
	for i := 0; i < len(names); i++ {
		println(names[i])
	}

	// := range similar to using enumerate in python
	for i, name := range names { // i is the index and name is the value
		println(i, name)
	}

	// if you want to ignore the index, you can use the _ (underscore) character
	for _, name := range names { // _ is used to ignore the index
		println(name)
	}



}

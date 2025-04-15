package main

import (
	"fmt"
	"math"
	"strings"
)

func Greeting(name string) {
	fmt.Println("Hello, " + name + "!")
	return // this is optional

}

func cycle(name []string, f func(string)) {
	for _, v := range name {
		f(v)
	}
}

func CircleArea(r float64) float64{

	area := math.Pi * r * r
	// fmt.Println("Area of circle is ", area)
	return area
}

//return multiple values
func CircleFeatures(r float64) (float64, float64){
	area := math.Pi * r * r
	circumference := 2 * math.Pi * r
	return area, circumference
}

func GetInitials(name string) string {
	initials := ""
	names := strings.Split(strings.ToUpper(name), " ")
	for _, v := range names {
		initials += string(v[0])
	}
	return initials
}

func main(){
	// Greeting("Derek")
	cycle([]string{"Derek", "Amuna", "John"}, Greeting)
	r := 22.0
	area := CircleArea(r)

	fmt.Printf("Area of the circle with radius %.2f is %.2f\n", r, area)

	area, circumference := CircleFeatures(r)
	fmt.Printf("A circle with radius %.2f has area %.2f and circumference %.2f\n", r, area, circumference)
	name := "Derek nsoh Amuna"
	inits := GetInitials(name)
	fmt.Printf("Initials of %s is %s\n", name,inits)
}

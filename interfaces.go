package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type shape interface{
	area() float64
	perimeter() float64
}

type square struct{
	length float64
}

type circle struct{
	radius float64
}

func (s square) area() float64{
	return s.length * s.length
}

func (s square) perimeter() float64{
	return 4 * s.length
}

func (c circle) area() float64{
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64{
	return 2 * math.Pi * c.radius
}

func ShapeInfo(s shape) (float64, float64){

	return s.area(), s.perimeter()
}

func getInput(prompt string) (string, error){
	reader := bufio.NewReader(os.Stdin)
	var input string
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

func getFloatInput()(float64){
	input, err := getInput("Enter a number: ")
	if err != nil {
		return 0
	}

	f, err1 := strconv.ParseFloat(input, 64)

	if err1 != nil {
		return 0
	}
	return f
}

func main(){
	sq := square{length: getFloatInput()}
	cir := circle{radius: getFloatInput()}

	sqArea, sqPerimeter := ShapeInfo(sq)
	cirArea, cirPerimeter := ShapeInfo(cir)

	fmt.Println("Square Area:", sqArea, "Square Perimeter:", sqPerimeter)
	fmt.Println("Circle Area:", cirArea, "Circle Perimeter:", cirPerimeter)
}

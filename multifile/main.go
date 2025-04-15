package main

import "fmt"

func main(){
	fmt.Println(points)
	points = append(points, 100)
	total, average := PointsSummary(points)
	fmt.Println("Total:", total)
	fmt.Println("Average:", average)
}

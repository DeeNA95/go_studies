package main
// all of these are auto imported to any other file in the same package (ie main)
import (
	"gonum.org/v1/gonum/stat"
	"github.com/thoas/go-funk"
)

var points = []float64{10,92,73,20,45,67,89,12,34,56}

func PointsSummary(points []float64) (int, float64) {

	total := funk.Sum(points)

	average := stat.Mean(points, nil)

	return int(total), average
}

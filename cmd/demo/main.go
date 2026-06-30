package main

import (
	"fmt"

	"github.com/NickolasBlack/finmath/math"
)

func main() {
	// Initialize test point
	testPoint := math.Point{
		X: 3,
		Y: 4,
	}
	otherTestPoint := math.Point{
		X: 2,
		Y: 2,
	}
	fmt.Println(testPoint.Equals(testPoint))
	fmt.Println(testPoint.Equals(otherTestPoint))
}

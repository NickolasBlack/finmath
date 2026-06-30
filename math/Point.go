package math

import (
	"fmt"
	"math"
)

/*
	Point Class
	contains and x and y value
*/

type Point struct {
	X float64 // x value
	Y float64 // y value
}

/* Point Struct computational series */

// Distance methods
func (pt *Point) Distance() float64 {
	// Find the distance between a point and the origin
	distanceX := pt.X * pt.X                // (x - 0)^2
	distanceY := pt.Y * pt.Y                // (y - 0)^2
	return math.Sqrt(distanceX + distanceY) // sqrt(x + y)
}

func (pt *Point) DistanceTo(otherPoint Point) float64 {
	// Find the distance to another point
	distanceX := math.Pow(pt.X-otherPoint.X, 2)
	distanceY := math.Pow(pt.Y-otherPoint.Y, 2)
	return math.Sqrt(distanceX + distanceY)
}

// Addition methods
func (pt *Point) Add(factor float64) {
	// Add a number to a point
	pt.X = pt.X + factor
	pt.Y = pt.Y + factor
}
func (pt *Point) AddTo(otherPoint Point) Point {
	// Adds point to another point. returns a point
	result := Point{
		// initialize new point
		pt.X + otherPoint.X, // add x
		pt.Y + otherPoint.Y, // add y
	}
	return result // return result
}

// Subtraction series
func (pt *Point) Subtract(factor float64) {
	// Subtract number from point
	pt.X = pt.X - factor // subtract x
	pt.Y = pt.Y - factor // subtract y
}

func (pt *Point) SubtractPoint(otherPoint Point) Point {
	// Subtract a point from a point
	result := Point{
		// Initialize new point
		pt.X - otherPoint.X, // subtract x
		pt.Y - otherPoint.Y, // subtract y
	}
	return result // return result
}

// Scalar series
func (pt *Point) Scale(scalar float64) {
	// Multiply by a scalar
	pt.X = pt.X * scalar // scale x
	pt.Y = pt.Y * scalar // scale y
}
func (pt *Point) ScalePoint(otherPoint Point) Point {
	// Multiplies 2 point together
	result := Point{
		// Initialize new point
		pt.X * otherPoint.X, // scale x
		pt.Y * otherPoint.Y, // scale y
	}
	return result // return result
}

// Point Struct formatter and utility series
func (pt Point) String() string {
	// Return a fomatted string
	return fmt.Sprintf("(%.2f, %.2f)", pt.X, pt.Y)
}
func (pt *Point) Equals(otherPoint Point) bool {
	// Checks if two points are equal to each other
	return pt.X == otherPoint.X && pt.Y == otherPoint.Y
}

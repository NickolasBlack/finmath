package math

import (
	"math"
	"testing"
)

/*
	Unit Test for Point Structs
*/

func TestPointDistance(t *testing.T) {
	// Point distance unit test
	tests := []struct {
		// Define testing table
		name     string
		point    Point
		expected float64
	}{
		{
			// Test Origin
			name:     "Origin",
			point:    Point{0, 0},
			expected: 0,
		},
		{
			name:     "3-4-5 Triangle",
			point:    Point{3, 4},
			expected: 5,
		},
		{
			name:     "Negative Coordinates",
			point:    Point{-3, -4},
			expected: 5,
		},
		{
			name:     "Decimal Coordinates",
			point:    Point{1.5, 2.5},
			expected: math.Sqrt(8.5),
		},
	}

	// Initialize tolerance
	const epsilon = 1e-9

	// run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := test.point.Distance()

			if math.Abs(actual-test.expected) > epsilon {
				t.Errorf("Distance() = %f, want %f", actual, test.expected)
			}
		})
	}
}

func TestDistanceTo(t *testing.T) {
	// Test Distance to
	tests := []struct {
		// define the test table
		name     string
		point1   Point
		point2   Point
		expected float64
	}{
		{
			name:     "Same Point",
			point1:   Point{0, 0},
			point2:   Point{0, 0},
			expected: 0,
		},
		{
			name:     "3-4-5 Triangle",
			point1:   Point{0, 0},
			point2:   Point{3, 4},
			expected: 5,
		},
		{
			name:     "Negative Coordinates",
			point1:   Point{-3, -4},
			point2:   Point{0, 0},
			expected: 5,
		},
		{
			name:     "Diagonal",
			point1:   Point{1, 1},
			point2:   Point{4, 5},
			expected: 5,
		},
		{
			name:     "Decimal Coordinates",
			point1:   Point{1.5, 2.5},
			point2:   Point{4.5, 6.5},
			expected: 5,
		},
	}

	// define tolerance
	const epsilon = 1e-9

	// run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := test.point1.DistanceTo(test.point2)

			if math.Abs(actual-test.expected) > epsilon {
				t.Errorf("DistanceTo() = %f, want %f", actual, test.expected)
			}
		})
	}

}

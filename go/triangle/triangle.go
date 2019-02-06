// Package triangle determines if a triangle is equilateral, isosceles, or scalene.
package triangle

import (
	"math"
)

// Kind is of type string
type Kind string

const (
	// NaT "not a triangle"
	NaT = "not a triangle"
	// Equ "equilateral"
	Equ = "equilateral"
	// Iso "isoceles"
	Iso = "isosceles"
	// Sca "scalene"
	Sca = "scalene"
)

// KindFromSides determines the kind of triangle from the length of its three sides.
func KindFromSides(a, b, c float64) Kind {
	areNumbers := func(a, b, c float64) bool {
		return !math.IsNaN(a) || !math.IsNaN(b) || !math.IsNaN(c)
	}

	areInfinity := func(a, b, c float64) bool {
		return math.IsInf(a, 1) || math.IsInf(b, 1) || math.IsInf(c, 1) ||
			math.IsInf(a, -1) || math.IsInf(b, -1) || math.IsInf(c, -1) ||
			math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0)
	}

	allSidesPositive := func(a, b, c float64) bool {
		if !areNumbers(a, b, c) || areInfinity(a, b, c) {
			return false
		}

		return a > 0 && b > 0 && c > 0
	}

	triangleEquality := func(a, b, c float64) bool {
		return a+b >= c && b+c >= a && a+c >= b
	}

	isEquilateral := func(a, b, c float64) bool {
		return a == b && b == c && a == c
	}

	isIsosceles := func(a, b, c float64) bool {
		return a == b || b == c || a == c
	}

	isScalene := func(a, b, c float64) bool {
		return a != b && b != c && a != c
	}

	isNotATriangle := func(a, b, c float64) bool {
		return !allSidesPositive(a, b, c) || !triangleEquality(a, b, c)
	}

	switch {
	case isNotATriangle(a, b, c):
		return NaT
	case isEquilateral(a, b, c):
		return Equ
	case isIsosceles(a, b, c):
		return Iso
	case isScalene(a, b, c):
		return Sca
	default:
		return NaT
	}
}

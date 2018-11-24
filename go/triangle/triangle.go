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
	var k Kind

	if isNotATriangle(a, b, c) {
		k = NaT
	} else if isEquilateral(a, b, c) {
		k = Equ
	} else if isIsosceles(a, b, c) {
		k = Iso
	} else if isScalene(a, b, c) {
		k = Sca
	}

	return k
}

func isNotATriangle(a, b, c float64) bool {
	return !allSidesPositive(a, b, c) || !triangleEquality(a, b, c)
}

func allSidesPositive(a, b, c float64) bool {
	if !areNumbers(a, b, c) || areInfinity(a, b, c) {
		return false
	}

	return a > 0 && b > 0 && c > 0
}

func areNumbers(a, b, c float64) bool {
	return !math.IsNaN(a) || !math.IsNaN(b) || !math.IsNaN(c)
}

func areInfinity(a, b, c float64) bool {
	return math.IsInf(a, 1) || math.IsInf(b, 1) || math.IsInf(c, 1) ||
		math.IsInf(a, -1) || math.IsInf(b, -1) || math.IsInf(c, -1) ||
		math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0)
}

func triangleEquality(a, b, c float64) bool {
	return a+b >= c && b+c >= a && a+c >= b
}

func isEquilateral(a, b, c float64) bool {
	return a == b && b == c && a == c
}

func isIsosceles(a, b, c float64) bool {
	return a == b || b == c || a == c
}

func isScalene(a, b, c float64) bool {
	return a != b && b != c && a != c
}

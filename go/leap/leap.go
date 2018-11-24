// Package leap reports if a given year is a leap year
package leap

// IsLeapYear returns true if given year is a leap year, false otherwise
func IsLeapYear(year int) bool {
	if divisibleBy(year, 4) {
		if divisibleBy(year, 100) {
			if divisibleBy(year, 400) {
				return true
			}
			return false
		}
		return true
	}
	return false
}

func divisibleBy(year int, divisor int) bool {
	return year%divisor == 0
}

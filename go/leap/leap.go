// Package leap reports if a given year is a leap year
package leap

// IsLeapYear returns true if given year is a leap year, false otherwise
func IsLeapYear(year int) bool {

	if year%4 != 0 {
		return false
	}
	if year%400 == 0 {
		return true
	}
	if year%100 == 0 {
		return false
	}

	return true
}

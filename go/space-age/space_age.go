// Package space calculates how old someone would be on a given planet by their
// age in earth seconds
package space

// Planet is a type of string
type Planet string

// Age converts earth seconds to given planet's years
func Age(seconds float64, planet Planet) float64 {
	var base float64 = 31557600

	age := func(year float64) float64 {
		return seconds / (base * year)
	}

	switch planet {
	case "Earth":
		return age(1)
	case "Mercury":
		return age(0.2408467)
	case "Venus":
		return age(0.61519726)
	case "Mars":
		return age(1.8808158)
	case "Jupiter":
		return age(11.862615)
	case "Saturn":
		return age(29.447498)
	case "Uranus":
		return age(84.016846)
	case "Neptune":
		return age(164.79132)
	default:
		return age(1)
	}
}

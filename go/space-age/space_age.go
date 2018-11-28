// Package space calculates how old someone would be on a given planet by their
// age in earth seconds
package space

// Planet is a type of string
type Planet string

// Age converts earth seconds to given planet's years
func Age(seconds float64, planet Planet) float64 {
	return float64(seconds) / YearSeconds(planet)
}

// YearSeconds returns a given planet's year in seconds
func YearSeconds(planet Planet) float64 {
	var base float64 = 31557600

	planetYears := func(planet Planet) float64 {
		switch planet {
		case "Earth":
			return 1
		case "Mercury":
			return 0.2408467
		case "Venus":
			return 0.61519726
		case "Mars":
			return 1.8808158
		case "Jupiter":
			return 11.862615
		case "Saturn":
			return 29.447498
		case "Uranus":
			return 84.016846
		case "Neptune":
			return 164.79132
		default:
			return 1
		}
	}

	return base * planetYears(planet)
}

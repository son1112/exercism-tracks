// Package space - given an age in seconds, calculate how old someone would be on
// other planets
package space

type Planet string

// Age: Float, String => Integer
// Takes age in seconds and planet name
// Returns age in planet's years
func Age(seconds float64, planet Planet) float64 {
	return float64(seconds) / YearSeconds(planet)
}

// YearSeconds: String => Float
// Takes planet name
// Returns planet year in seconds
func YearSeconds(planet Planet) float64 {
	var base float64 = 31557600

	var planetYears = map[Planet]float64{
		"Earth":   1,
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}

	return base * planetYears[planet]
}

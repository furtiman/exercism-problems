// Package space contains solution and tests for "space_age" exercise
package space

// Planet type implements string to represent a certain planet
type Planet string

// Age returns the age of the person on certain planet relative to earth seconds
func Age(seconds float64, planet Planet) float64 {
	switch planet {
	case "Earth":
		return seconds / 31557600
	case "Mercury":
		return seconds / 7600543.82
	case "Venus":
		return seconds / 19414149.1
	case "Mars":
		return seconds / 59354032.7
	case "Jupiter":
		return seconds / 374355659
	case "Saturn":
		return seconds / 929292363
	case "Uranus":
		return seconds / 2651370019.3296
	case "Neptune":
		return seconds / 5200418560.032
	}
	return 0
}

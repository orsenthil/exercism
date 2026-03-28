package space

type Planet string

var earthSeconds float64 = 31557600

func Age(seconds float64, planet Planet) float64 {
	if planet == "Earth" {
		return seconds / earthSeconds
	} else if planet == "Mercury" {
		return seconds / (earthSeconds * 0.2408467)
	} else if planet == "Venus" {
		return seconds / (earthSeconds * 0.61519726)
	} else if planet == "Mars" {
		return seconds / (earthSeconds * 1.8808158)
	} else if planet == "Jupiter" {
		return seconds / (earthSeconds * 11.862615)
	} else if planet == "Saturn" {
		return seconds / (earthSeconds * 29.447498)
	} else if planet == "Uranus" {
		return seconds / (earthSeconds * 84.016846)
	} else if planet == "Neptune" {
		return seconds / (earthSeconds * 164.79132)
	}
	return -1
}

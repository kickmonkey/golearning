package space

type Planet string

func Age(seconds float64, planet Planet) float64 {
    switch planet{
        case "Mercury":
        	return seconds/0.2408467/31557600.0
        case "Venus":
        	return seconds/0.61519726/31557600.0
        case "Earth":
        	return seconds/1.0/31557600.0
        case "Mars":
        	return seconds/1.8808158/31557600.0
        case "Jupiter":
        	return seconds/11.862615/31557600.0
        case "Saturn":
        	return seconds/29.447498/31557600.0
        case "Uranus":
        	return seconds/84.016846/31557600.0
        case "Neptune":
        	return seconds/164.79132/31557600.0
    }
    return -1.00
	panic("Please implement the Age function")
}

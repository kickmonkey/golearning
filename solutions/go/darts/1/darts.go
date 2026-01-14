package darts

func Score(x, y float64) int {
    a := x*x + y*y
    switch {
        case a <= 1.0 && a >= 0.0:
        	return 10
        case a > 1.0 && a <= 25.0:
        	return 5
        case a > 25.0 && a <= 100.0:
        	return 1
    }
	return 0
}

package triangle

// Kind represents the type of triangle.
type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// KindFromSides determines the kind of triangle based on side lengths.
func KindFromSides(a, b, c float64) Kind {
	// Sides must be positive
	if a <= 0 || b <= 0 || c <= 0 {
		return NaT
	}

	// Triangle inequality
	if a+b < c || a+c < b || b+c < a {
		return NaT
	}

	// All sides equal
	if a == b && b == c {
		return Equ
	}

	// Two sides equal
	if a == b || b == c || a == c {
		return Iso
	}

	// All sides different
	return Sca
}

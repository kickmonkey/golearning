package resistorcolor
import "strings"
// Colors returns the list of all colors.
func Colors() []string {
    return []string{"black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"}
	panic("Please implement the Colors function")
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
    colors := Colors()
    for i,v := range colors{
        if strings.ToLower(color) == v{
            return i
        }
    }
    return 11
}

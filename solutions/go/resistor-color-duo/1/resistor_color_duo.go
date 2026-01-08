package resistorcolorduo

func Value(colors []string) int {
    code := map[string]int{
        "black":  0,
        "brown":  1,
        "red":    2,
        "orange": 3,
        "yellow": 4,
        "green":  5,
        "blue":   6,
        "violet": 7,
        "grey":   8,
        "white":  9,
    }

    return code[colors[0]]*10 + code[colors[1]]
}

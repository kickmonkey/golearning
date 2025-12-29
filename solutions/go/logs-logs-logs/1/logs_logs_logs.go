package logs

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, r := range log {
		switch r {
		case '❗':
			return "recommendation"
		case '🔍': // U+1F50D
			return "search"
		case '☀': // U+2600
			return "weather"
		}
	}
    return "default"
	panic("Please implement the Application() function")
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
    rs := []rune(log)
    for i, r := range rs {
        if r == oldRune{
            rs[i] = newRune
        }
    }
    return string(rs)
	panic("Please implement the Replace() function")
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
    if limit < 0 {
		return false
	}

	count := 0
	for range log { // rune-based iteration
		count++
		if count > limit {
			return false
		}
	}
	return true
	panic("Please implement the WithinLimit() function")
}

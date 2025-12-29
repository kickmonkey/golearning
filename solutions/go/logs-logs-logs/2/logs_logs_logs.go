package logs

// Application identifies the application emitting the given log.
func Application(log string) string {
    for _, v := range log{
        switch v{
            case '❗':
            	return "recommendation"
            case '🔍':
            	return "search"
            case '☀':
            	return "weather"
        }
    }
    return "default"
	panic("Please implement the Application() function")
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
    r := []rune(log)
    for i, v := range r{
        if v == oldRune{
            r[i] = newRune
        }
    }
    return string(r)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
    if limit < 0 {
        return false
    }
    count := 0
    for range log{
        count++
        if count > limit{
            return false
        }
    }
    return true
}

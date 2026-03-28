package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
    for _, c := range log {
        if c == '❗' {
            return "recommendation"
        } else if c == '🔍' {
        	return "search"
        } else if c == '☀' {
        	return "weather"
        }
    }
    return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
    var newlog string
    for _, c := range log {
        if c == oldRune {
            newlog += string(newRune)
        } else {
        	newlog += string(c)
        }
    }
	return newlog
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
    if utf8.RuneCountInString(log) <= limit {
        return true
    }
	return false
}

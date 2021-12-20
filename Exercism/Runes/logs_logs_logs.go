package logs

import (
	"strings"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, v := range log {
		if v == '❗' {
			return "recommendation"
		}
		if v == '🔍' {
			return "search"
		}
		if v == '☀' {
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	for _, v := range log {
		if v != newRune {
			log = strings.ReplaceAll(log, string(oldRune), string(newRune))
			return log
		}
	}
	return log
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	for range log {
		if utf8.RuneCountInString(log) > limit {
			return false
		}
	}
	return true
}

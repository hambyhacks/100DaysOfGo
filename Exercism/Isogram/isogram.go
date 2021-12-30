package isogram

import (
	"strings"
)

func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	wordRune := []rune(word)

	for i, j := range word {
		charCount := strings.Count(word, string(j))

		if wordRune[i] == '-' || wordRune[i] == ' ' {
			continue
		}

		if charCount > 1 {
			return false
		}

	}
	return true
}

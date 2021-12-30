package scrabble

import "strings"

func Score(word string) int {
	word = strings.ToUpper(word)
	wordRune := []rune(word)

	var score int
	for i := range wordRune {
		switch wordRune[i] {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			score += 1
			continue
		case 'D', 'G':
			score += 2
			continue
		case 'B', 'C', 'M', 'P':
			score += 3
			continue
		case 'F', 'H', 'V', 'W', 'Y':
			score += 4
			continue
		case 'K':
			score += 5
			continue
		case 'J', 'X':
			score += 8
			continue
		case 'Q', 'Z':
			score += 10
			continue
		}
	}
	return score
}

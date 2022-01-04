package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

func Valid(id string) bool {
	checksum := 0
	id = strings.ReplaceAll(id, " ", "")
	parity := len(id) % 2

	for i, j := range id {
		if !unicode.IsDigit(j) {
			return false
		}

		idDigits, _ := strconv.Atoi(string(j))
		if i%2 == parity {
			idDigits *= 2
			if idDigits > 9 {
				idDigits -= 9
			}
		}
		checksum += idDigits
	}
	return checksum%10 == 0 && len(id) > 1
}

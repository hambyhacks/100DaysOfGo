package raindrops

import "strconv"

func Convert(number int) string {
	if number%3 == 0 && number%5 == 0 && number%7 == 0 {
		return "PlingPlangPlong"
	}

	if number%3 == 0 && number%5 == 0 {
		return "PlingPlang"
	}

	if number%5 == 0 && number%7 == 0 {
		return "PlangPlong"
	}

	if number%3 == 0 && number%7 == 0 {
		return "PlingPlong"
	}

	if number%3 == 0 {
		return "Pling"
	}

	if number%5 == 0 {
		return "Plang"
	}

	if number%7 == 0 {
		return "Plong"
	}
	return strconv.Itoa(number)
}

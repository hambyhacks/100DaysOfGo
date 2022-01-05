package grains

import (
	"errors"
	"math"
)

func Square(number int) (uint64, error) {
	squares := 1 * int(math.Pow(2, float64(number)-1))

	if number <= 0 || number > 64 {
		return 1, errors.New("invalid square")
	}

	return uint64(squares), nil
}

func Total() uint64 {
	return 1<<64 - 1
}

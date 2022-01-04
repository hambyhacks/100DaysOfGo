package letter

type FreqMap map[rune]int

// Line 19-21: Anonymous function call.

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(l []string) FreqMap {
	ch := make(chan FreqMap)
	n := FreqMap{}

	for _, j := range l {
		go func(j string) {
			ch <- Frequency(j)
		}(j)

		for x, y := range <-ch {
			n[x] += y
		}
	}
	return n
}

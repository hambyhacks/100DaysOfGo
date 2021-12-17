package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard contains a map of eight Ranks, accessed with values from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	count := 0
	for _, v := range cb[rank] {
		if !v {
			continue
		}
		count++
	}
	return count
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	if file > 8 || file < 1 {
		return 0
	}
	count := 0
	for _, values := range cb {
		if values[file-1] {
			count++
		}
	}
	return count
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	count := 0
	for range cb {
		count++
	}
	return count * 8
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	count := 0
	for i := range cb {
		count += CountInRank(cb, i)
	}
	return count
}

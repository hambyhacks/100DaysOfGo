package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {
	// Convert string to rune array.
	gene_A := []rune(a)
	gene_B := []rune(b)

	// Check if gene length is same.
	if len(gene_A) != len(gene_B) {
		return 0, errors.New("gene length mismatch")
	}
	// Calculate Hamming Distance
	distance := 0
	for i := range gene_A {
		if gene_B[i] != gene_A[i] {
			distance++
		}
	}
	return distance, nil
}

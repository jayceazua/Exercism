package hamming

import "errors"

// Credit: https://gobyexample.com/range
// Credit: https://golang.org/pkg/errors/
// Dylan Finn helped me get unblocked, with the error package.
func Distance(a, b string) /* return values -> */ (int, error) {
	differences := 0

	if len(a) != len(b) {
		return differences, errors.New("...")
	}
	// for loop to count the differences
	for i := range a {
		if a[i] != b[i] {
			differences += 1
		}
	}

	return differences, nil
}

/*
By counting the number of differences between two homologous DNA strands taken from different genomes with a common ancestor,
we get a measure of the minimum number of point mutations that could have occurred on the evolutionary path between the two strands.

This is called the 'Hamming distance'.

It is found by comparing two DNA strands and counting how many of the nucleotides are different from their equivalent in the other string

*/

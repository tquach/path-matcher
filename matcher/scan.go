package matcher

import (
	"bufio"
	"io"
	"strconv"
)

type PatternsAndPaths []string

// Scan in the input and produce a slice of PatternsAndPaths and a possibly nil error.
func Scan(r io.Reader) ([]PatternsAndPaths, error) {
	// Initialize a new scanner using the given reader
	scanner := bufio.NewScanner(r)

	var results []PatternsAndPaths
	for scanner.Scan() {
		txt := scanner.Text()

		// First line is the number of rows
		if n, err := strconv.Atoi(txt); err == nil {
			var rows []string

			// Scan in the n rows containing either patterns
			// or paths (to be determined later)
			for len(rows) < n && scanner.Scan() {
				rows = append(rows, scanner.Text())
			}
			results = append(results, PatternsAndPaths(rows))
		}
	}

	// Check for any errors
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

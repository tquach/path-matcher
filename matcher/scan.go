package matcher

import (
	"bufio"
	"io"
	"strconv"
)

type PatternsAndPaths []string

func Scan(r io.Reader) ([]PatternsAndPaths, error) {
	scanner := bufio.NewScanner(r)
	var results []PatternsAndPaths
	for scanner.Scan() {
		txt := scanner.Text()

		// First line is the number of rows
		if n, err := strconv.Atoi(txt); err == nil {
			var rows []string
			for len(rows) < n && scanner.Scan() {
				rows = append(rows, scanner.Text())
			}
			results = append(results, PatternsAndPaths(rows))
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

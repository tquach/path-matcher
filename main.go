package main

import (
	"log"
	"os"

	"github.com/tantastik/path-matcher/matcher"
)

var logger *log.Logger

func main() {
	patternsAndPaths, err := matcher.Scan(os.Stdin)
	if err != nil {
		logger.Fatalln("Failed to read standard in.")
	}

	if len(patternsAndPaths) != 2 {
		logger.Fatalln("Invalid format of input.", patternsAndPaths)
	}

	patterns := patternsAndPaths[0]
	paths := patternsAndPaths[1]

	m := matcher.NewPatternPathMatcher(patterns)
	for _, path := range paths {
		pattern := m.BestMatch(path)
		logger.Println(pattern)
	}
}

func init() {
	logger = log.New(os.Stdout, "", 0)
}

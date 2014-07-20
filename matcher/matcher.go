// Package matcher contains the matching algorithm
package matcher

import "strings"

type Pattern struct {
	Fields []string
	Raw    string
}

func (p *Pattern) Matches(str string) bool {
	// Trim the path of leading / trailing slashes and spaces
	elems := strings.Split(strings.Trim(str, "/ "), "/")
	if len(elems) != len(p.Fields) {
		return false
	}

	for i := 0; i < len(elems); i++ {
		expr, val := p.Fields[i], elems[i]
		if val == "" {
			return false
		}

		switch expr {
		case val:
		case "*":
			continue
		default:
			return false
		}
	}
	return true
}

// Define a struct to hold our list of patterns
type PatternPathMatcher struct {
	// A list of patterns
	Patterns []Pattern
}

func NewPattern(pattern string) Pattern {
	return Pattern{strings.Split(pattern, ","), pattern}
}

func NewPatternPathMatcher(patternList []string) PatternPathMatcher {
	var patterns []Pattern

	for _, pattern := range patternList {
		patterns = append(patterns, NewPattern(pattern))
	}
	return PatternPathMatcher{patterns}
}

// Find the best matching pattern given the path.
func (m *PatternPathMatcher) BestMatch(str string) string {
	var matches []Pattern
	for _, pattern := range m.Patterns {
		// If there is an error, the value of matched will be false
		if pattern.Matches(str) {
			matches = append(matches, pattern)
		}
	}

	// Check for ties
	if len(matches) > 0 {
		OrderedBy(wildcards, rightMostWildcard).Sort(matches)
		return matches[0].Raw
	} else {
		return "NO MATCH"
	}
}

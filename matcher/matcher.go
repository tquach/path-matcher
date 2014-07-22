package matcher

import "strings"

type Pattern struct {
	Fields []string
	Raw    string
}

// Matches a given string against itself. This algorithm is simple and straightforward
// using an exact match or a single wildcard match.
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

		// Naive implementation for matching
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

// Constructor factory for creating a new Pattern instance. Also splits the pattern against any fields
func NewPattern(pattern string) Pattern {
	return Pattern{strings.Split(pattern, ","), pattern}
}

// Create a new PatternPathMatcher instance with a slice of patterns
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

	// Calculate all matches
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

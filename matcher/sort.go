// This implementation is based off Go standard library multi-key sort example.
package matcher

import (
	"sort"
	"strings"
)

// A function definition type for determining how to sort
type sortFn func(p1, p2 *Pattern) bool

// A struct that will implement the sort.Interface
type patternSorter struct {
	patterns []Pattern
	sortFns  []sortFn
}

// Len is part of sort.Interface.
func (s *patternSorter) Len() int {
	return len(s.patterns)
}

// Swap is part of sort.Interface.
func (s *patternSorter) Swap(i, j int) {
	s.patterns[i], s.patterns[j] = s.patterns[j], s.patterns[i]
}

// Less is part of sort.Interface. It is implemented by calling the "sortFn" closure in the sorter.
func (s *patternSorter) Less(i, j int) bool {
	p, q := &s.patterns[i], &s.patterns[j]
	// Try all but the last comparison.
	var k int
	for k = 0; k < len(s.sortFns)-1; k++ {
		sortBy := s.sortFns[k]
		switch {
		case sortBy(p, q):
			return true
		case sortBy(q, p):
			return false
		}
	}
	return s.sortFns[k](p, q)
}

// Sort a slice of Pattern objects. Delegates to the sort.Sort function.
func (as *patternSorter) Sort(patterns []Pattern) {
	as.patterns = patterns
	sort.Sort(as)
}

// A function that applies a variadic list of closures and returns a sorter.
func OrderedBy(sortFn ...sortFn) *patternSorter {
	return &patternSorter{
		sortFns: sortFn,
	}
}

// Closures used for sorting lists of Patterns
var wildcards = func(p1, p2 *Pattern) bool {
	return strings.Count(p1.Raw, "*") < strings.Count(p2.Raw, "*")
}

// Sort by the right-most wildcard using the strings.Index function
var rightMostWildcard = func(p1, p2 *Pattern) bool {
	return strings.Index(p1.Raw, "*") > strings.Index(p2.Raw, "*")
}

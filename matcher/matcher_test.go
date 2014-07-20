package matcher

import . "gopkg.in/check.v1"

type MatcherTestSuite struct{}

var _ = Suite(&MatcherTestSuite{})

var testPatterns = []string{
	"*,b,*",
	"a,*,*",
	"*,*,c",
	"foo,bar,baz",
	"w,x,*,*",
	"*,x,y,z",
}

var testPaths = []string{
	"/w/x/y/z/",
	"a/b/c",
	"foo/",
	"foo/bar/",
	"foo/bar/baz/",
}

var expected = []string{
	"*,x,y,z",
	"a,*,*",
	"NO MATCH",
	"NO MATCH",
	"foo,bar,baz",
}

func (suite *MatcherTestSuite) TestPatternMatchesBasicWildcardPattern(c *C) {
	pattern := NewPattern("*,b,*")
	matches := pattern.Matches("alongstring/b/abskjsdfkjsfc")
	c.Check(matches, Equals, true)

	matches = pattern.Matches("/b/")
	c.Check(matches, Equals, false)

	matches = pattern.Matches("a/b/d")
	c.Check(matches, Equals, true)

	matches = pattern.Matches("a/b/c/b/p")
	c.Check(matches, Equals, false)

	matches = pattern.Matches("*/b/*")
	c.Check(matches, Equals, true)

	matches = pattern.Matches("b")
	c.Check(matches, Equals, false)
}

func (suite *MatcherTestSuite) TestFieldLengthMatching(c *C) {
	pattern := NewPattern("foo,bar,baz")

	matches := pattern.Matches("foo/")
	c.Check(matches, Equals, false)

	matches = pattern.Matches("foo/bar/")
	c.Check(matches, Equals, false)

	matches = pattern.Matches("foo/bar/baz/")
	c.Check(matches, Equals, true)
}

func (suite *MatcherTestSuite) TestMatchAllPatterns(c *C) {
	pathMatcher := NewPatternPathMatcher(testPatterns)

	for i, testPath := range testPaths {
		obtained := pathMatcher.BestMatch(testPath)
		c.Check(obtained, Equals, expected[i])
	}
}

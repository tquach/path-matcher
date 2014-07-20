package matcher

import (
	"bytes"
	. "gopkg.in/check.v1"
)

type ScannerTestSuite struct{}

var _ = Suite(&ScannerTestSuite{})

var testdata = []string{
	"0\n0",
	"1\na,*,c\n0",
	"1\na,b,*\n1\na/b/d",
}

func (suite *ScannerTestSuite) TestScanInputFile(c *C) {
	for _, test := range testdata {
		buf := bytes.NewBufferString(test)
		patternspaths, err := Scan(buf)

		c.Assert(err, IsNil)
		c.Check(len(patternspaths), Equals, 2)
	}
}

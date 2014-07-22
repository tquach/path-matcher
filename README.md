## Path Matcher

[![Build Status](https://travis-ci.org/tantastik/path-matcher.png?branch=master)](https://travis-ci.org/tantastik/path-matcher)

Simple example of using Go's Scanner library to scan a file and process it. 

## Usage

```go
go get -u ./...
go build
./path-matcher < test/testfile1.txt
```

Outputs results to standard out.

## Documentation

Documentation is done through [go docs](http://godoc.org/code.google.com/p/go.tools/cmd/godoc):

```
godoc -http=:6060
```

Open a browser to http://localhost:6060/pkg/github.com/tantastik/path-matcher/matcher/

## Analysis

Basic algorithm is O(mn) in runtime, where m is the number of paths and n is the number of patterns:

	for each path in m paths
		for each pattern in n patterns
			check for match

To check if it's a match, it's simple enough to compare for exact match or presence of a wildcard, so this can be done in constant time.

To find the best match, we keep all matches in an array and sort the array based on the "score". The score is calculated using least number of wildcards and by the furthest left wildcard that appears furthest to the right in the pattern. The lowest score is the winner. The sorting algorithm is quickSort and can be done in logarithmic time.

## Improvements

Can we do better than O(mn)? Some ideas for improvements include doing preprocessing on the lists beforehand, as suggested by Boyer-Moore string matching algorithm. This allows us to disregard any remaining patterns when a bad character is encountered. This will give us O(m+n) for worst case scenario.

We can also group together patterns into buckets and when matching paths, we can know immediately if there is a match if there is no bucket that corresponds to the first element in the path.

Most of the improvements require some kind of preprocessing and shortcut to avoid having to do further comparisons.

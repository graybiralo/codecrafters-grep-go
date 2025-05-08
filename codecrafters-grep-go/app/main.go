package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) < 3 || os.Args[1] != "-E" {
		fmt.Fprintf(os.Stderr, "usage: mygrep -E <pattern>\n")
		os.Exit(2)
	}

	pattern := os.Args[2]

	line, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: reading input: %v\n", err)
		os.Exit(2)
	}

	match, err := matchLine(line, pattern)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(2)
	}

	if !match {
		fmt.Println("Match not found!")
		os.Exit(1)
	}
	fmt.Println("Match found!")
}

func matchLine(line []byte, pattern string) (bool, error) {
	// Regex to find backreferences like \1, \2
	backrefRegex := regexp.MustCompile(`\\(\d+)`)

	// Extract all backreference numbers
	backrefMatches := backrefRegex.FindAllStringSubmatch(pattern, -1)

	// Replace all backreferences like \1 with (.*)
	capturePattern := backrefRegex.ReplaceAllString(pattern, "(.*)")

	// new pattern
	re, err := regexp.Compile(capturePattern)
	if err != nil {
		return false, err
	}

	// Match against the input line
	matches := re.FindSubmatch(line)
	if matches == nil {
		return false, nil
	}

	// Compare captured groups to backreferences
	for _, match := range backrefMatches {
		// match[1] is the digit string (e.g., "1", "2")
		groupIndex := match[1]
		var index int
		fmt.Sscanf(groupIndex, "%d", &index)

		// index refers to the original backreference group
		// Example: for \1, check that group 1 == group N (where N is the corresponding (.*))
		// All (.*) captures go from matches[1] onward
		ref1 := matches[index]
		ref2 := matches[len(matches)-len(backrefMatches)+index-1]

		if !bytes.Equal(ref1, ref2) {
			return false, nil
		}
	}

	return true, nil
}

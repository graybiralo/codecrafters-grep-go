package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
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
	// Find all capturing group references like \1, \2, etc.
	backrefRegex := regexp.MustCompile(`\\(\d+)`)
	backrefs := backrefRegex.FindAllStringSubmatch(pattern, -1)

	// If there are no backreferences, use normal regex
	if len(backrefs) == 0 {
		re, err := regexp.Compile(pattern)
		if err != nil {
			return false, err
		}
		return re.Match(line), nil
	}

	// Replace \1, \2... with new capturing groups to match manually
	// Also escape literal backslashes
	capturePattern := backrefRegex.ReplaceAllString(pattern, "(.*)")

	re, err := regexp.Compile(capturePattern)
	if err != nil {
		return false, err
	}

	matches := re.FindSubmatch(line)
	if matches == nil {
		return false, nil
	}

	// Extract backreference positions and compare
	for _, backref := range backrefs {
		index := backref[1] // string number like "1", "2"
		groupNum := atoi(index)

		if groupNum >= len(matches) {
			return false, nil
		}

		// Compare the earlier group with the backreference location
		// Group numbers in regex are 1-indexed
		refValue := matches[groupNum]
		backrefGroup := matches[len(matches)-1] // last match is the backref we just captured
		if !bytes.Equal(refValue, backrefGroup) {
			return false, nil
		}
	}

	return true, nil
}

func atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

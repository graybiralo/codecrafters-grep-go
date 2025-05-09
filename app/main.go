package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
	"unicode/utf8"
)

// Ensures gofmt doesn't remove the "bytes" import above (feel free to remove this!)
var _ = bytes.ContainsAny

// Usage: echo <input_text> | your_program.sh -E <pattern>
func main() {
	if len(os.Args) < 3 || os.Args[1] != "-E" {
		fmt.Fprintf(os.Stderr, "usage: mygrep -E <pattern>\n")
		os.Exit(2) // 1 means no lines were selected, >1 means error
	}

	pattern := os.Args[2]

	line, err := io.ReadAll(os.Stdin) // assume we're only dealing with a single line
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: read input text: %v\n", err)
		os.Exit(2)
	}

	ok, err := matchLine(line, pattern)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(2)
	}

	if !ok {
		os.Exit(1)
	}

	// default exit code is 0 which means success
}

func matchLine(line []byte, pattern string) (bool, error) {
	switch pattern {
	case `\d`: // Matches any digit
		re := regexp.MustCompile(`\d`)
		return re.Match(line), nil
	case `\w`: // Matches any alphanumeric character or underscore
		re := regexp.MustCompile(`[a-zA-Z0-9_]`)
		return re.Match(line), nil
	}

	if utf8.RuneCountInString(pattern) != 1 {
		return false, fmt.Errorf("unsupported pattern: %q", pattern)
	}

	return bytes.Contains(line, []byte(pattern)), nil
}

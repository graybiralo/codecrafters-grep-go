# MyGrep - A Regex Matching Tool in Go

`mygrep` is a simplified, custom-built implementation of the Unix `grep` utility, written in Go. It supports regular expression matching with various features including character classes, anchors, quantifiers, and backreferences.

## Features

- Basic string matching with regular expressions
- Character classes: `\w`, `\d`, and custom sets like `[abc]` or `[^abc]`
- Anchors: `^` (start of line), `$` (end of line)
- Quantifiers: `+` (one or more)
- Capturing groups and backreferences:
  - Single backreference: `(cat) and \1`
  - Multiple backreferences: `(\d+) (\w+) squares and \1 \2 circles`
  - Nested backreferences: `('(cat) and \2') is the same as \1`

##  Usage

```bash
echo -n "<input>" | ./your_program.sh -E "<pattern>"

## Example

echo -n "apple" | ./your_program.sh -E "[^abc]"       # Match found!
echo -n "1 apple" | ./your_program.sh -E "\d apple"   # Match found!
echo -n "cat and cat" | ./your_program.sh -E "(cat) and \1"  # Match found!
echo -n "3 red squares and 3 red circles" | ./your_program.sh -E "(\d+) (\w+) squares and \1 \2 circles"  # Match found!




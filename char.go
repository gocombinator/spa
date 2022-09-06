package spa

import (
	"fmt"
	"unicode/utf8"

	"github.com/gocombinator/spa/internal"
)

// Char matches single rune based on provided pattern returning string.
//
// Example pattern:
//
//		Char(
//	   "0", "57",
//	   "begi", "xz"
//	  )
//
// Is interpreted as:
// - range "0"-"5" succeeds
// - single "7" succeeds
// - single "b" succeeds
// - single "e" succeeds
// - single "g" succeeds
// - range "i"-"x" succeeds
// - single "z" succeeds
//
// In other words:
// - listed runes in string match
// - last rune and first rune of next string pair - create range that match
// - ranges are between even and odd strings only
//
// As convention it's recommended to write pair of strings per line - as each line defines single range.
func Char(patterns ...string) Parser {

	// Map strings to runes outside of parser.
	var list [][]rune = make([][]rune, len(patterns))
	for i, s := range patterns {
		list[i] = []rune(s)
	}

	return func(in string) (any, int, error) {
		var r, _ = utf8.DecodeRuneInString(in)
		var prev rune
		for i, chars := range list {

			// Check range.
			if i%2 == 1 {
				if r >= prev && r <= chars[0] {
					var s = string(r)
					return s, len(s), nil
				}
			}

			// Check every char.
			for _, c := range chars {
				if r == c {
					var s = string(r)
					return s, len(s), nil
				}
			}

			prev = chars[len(chars)-1]
		}

		return nil, 0, fmt.Errorf("expected one of %s chars", internal.ReadablePatterns(patterns...))
	}
}

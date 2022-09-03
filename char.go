package spa

import "github.com/gocombinator/spa/internal"

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

	return func(in string) Result {
		for _, r := range in {
			var prev rune
			for i, chars := range list {

				// Check range.
				if i%2 == 1 {
					if r >= prev && r <= chars[0] {
						return Eat(in, len(string(r)))
					}
				}

				// Check every char.
				for _, c := range chars {
					if r == c {
						return Eat(in, len(string(r)))
					}
				}

				prev = chars[len(chars)-1]
			}

			//lint:ignore SA4004 we're range-ing to avoid whole input to rune conversion
			break
		}
		return Errorf(in, "expected one of %s chars", internal.ReadablePatterns(patterns...))
	}
}

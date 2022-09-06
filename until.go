package spa

import (
	"unicode/utf8"
)

// Until matches until provided parser succeeds or end of input is reached.
// Doesn't consume input of succeeding parser.
// Returns string with zero or more length.
//
//	(.*) (?p) -> $1
//
// Use [Min] to limit minimum length.
func Until(p Parser) Parser {
	return func(in string) (any, int, error) {
		var n = len(in)
		for o := 0; o < n; {
			if _, _, err := p(in[o:]); err != nil {
				var _, w = utf8.DecodeRuneInString(in[o:])
				o += w
			} else {
				return in[:o], o, nil
			}
		}
		return Rest(in)
	}
}

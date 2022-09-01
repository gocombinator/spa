package spa

import "github.com/gocombinator/spa/internal"

// Concat glues together results into single string.
func Concat(p Parser) Parser {
	return Map(p, internal.Concat)
}

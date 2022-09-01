package spa

import "github.com/gocombinator/spa/internal"

// Join glues together results into single string using provided separator.
func Join(p Parser, sep string) Parser {
	return Map(p, func(v any) string { return internal.Join(v, sep) })
}

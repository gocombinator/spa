package spa

import "github.com/gocombinator/spa/internal"

// Join glues together results into single string using provided separator.
func Join(sep string) func(Parser) Parser {
	return As(func(v any) string {
		return internal.Join(v, sep)
	})
}

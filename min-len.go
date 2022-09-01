package spa

import (
	"github.com/gocombinator/spa/internal"
)

// MinLen fails parser which produces result value with len less than provided value.
func MinLen(n int, p Parser) Parser {
	return func(in string) Result {
		if r := p(in); r.Error == nil {
			if l := internal.Len(r.Value); l < n {
				return Errorf("expected min len %d, got %d", n, l)
			} else {
				return r
			}
		} else {
			return r
		}
	}
}

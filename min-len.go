package spa

import (
	"fmt"

	"github.com/gocombinator/spa/internal"
)

// MinLen fails parser which produces result value with len less than provided value.
func MinLen(min int) Mapper {
	return func(p Parser) Parser {
		return func(in string) (any, int, error) {
			if v, n, err := p(in); err != nil {
				return nil, 0, err
			} else if l := internal.Len(v); l < min {
				return nil, 0, fmt.Errorf("expected min len %d, got %d", n, l)
			} else {
				return v, n, nil
			}
		}
	}
}

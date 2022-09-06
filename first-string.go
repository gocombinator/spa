package spa

import (
	"fmt"
	"strings"
)

// FirstString matches first string from given list.
// See [LongestString] for different strategy.
func FirstString(values []string) Parser {
	return func(in string) (any, int, error) {
		for _, value := range values {
			if strings.HasPrefix(in, value) {
				return value, len(value), nil
			}
		}
		return nil, 0, fmt.Errorf("expected first of %d strings", len(values))
	}
}

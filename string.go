package spa

import (
	"fmt"
	"sort"
	"strings"
)

// String matches longest string from given list.
// See [FirstString] for different strategy.
func String(values ...string) Parser {
	if len(values) == 1 {
		var value = values[0]
		return func(in string) (any, int, error) {
			if strings.HasPrefix(in, value) {
				return value, len(value), nil
			} else {
				return nil, 0, fmt.Errorf("expected %s string", value)
			}
		}
	}
	sort.Sort(sort.Reverse(sort.StringSlice(values)))
	return func(in string) (any, int, error) {
		for _, value := range values {
			if strings.HasPrefix(in, value) {
				return value, len(value), nil
			}
		}
		return nil, 0, fmt.Errorf("expected longest of %s strings", strings.Join(values, " / "))
	}
}

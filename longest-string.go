package spa

import (
	"sort"
	"strings"
)

// LongestString returns parser that matches longest string from given list.
// See [FirstString] for different strategy.
func LongestString(values []string) Parser {

	var sorted = make([]string, len(values))
	copy(sorted, values)
	sort.Sort(sort.Reverse(sort.StringSlice(sorted)))

	return func(in string) Result {
		for _, value := range sorted {
			if strings.HasPrefix(in, value) {
				return Ok(in[len(value):], value)
			}
		}
		return Errorf("expected longest of %d strings", len(sorted))
	}
}

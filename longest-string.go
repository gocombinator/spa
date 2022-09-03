package spa

import (
	"sort"
	"strings"
)

// LongestString matches longest string from given list.
// See [FirstString] for different strategy.
func LongestString(values []string) Parser {

	var sorted = make([]string, len(values))
	copy(sorted, values)
	sort.Sort(sort.Reverse(sort.StringSlice(sorted)))

	return func(in string) Result {
		for _, value := range sorted {
			if strings.HasPrefix(in, value) {
				return Eat(in, len(value))
			}
		}
		return Errorf(in, "expected longest of %s strings", strings.Join(sorted, " / "))
	}
}

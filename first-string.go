package spa

import "strings"

// FirstString matches first string from given list.
// See [LongestString] for different strategy.
func FirstString(values []string) Parser {
	return func(in string) Result {
		for _, value := range values {
			if strings.HasPrefix(in, value) {
				return Eat(in, len(value))
			}
		}
		return Errorf(in, "expected first of %d strings", len(values))
	}
}

package spa

import "strings"

// String matches given string.
func String(value string) Parser {
	return func(in string) Result {
		if strings.HasPrefix(in, value) {
			return Eat(in, len(value))
		} else {
			return Errorf(in, "expected %s string", value)
		}
	}
}

package spa

import "strings"

// String returns parser that matches given string.
func String(value string) Parser {
	return func(in string) Result {
		if strings.HasPrefix(in, value) {
			return Ok(in[len(value):], value)
		} else {
			return Errorf("expected " + value)
		}
	}
}

package spa

// Chars matches one or more chars with given patterns.
// See [Char] for patterns.
func Chars(patterns ...string) Parser {
	return Concat(Repeat(Char(patterns...)))
}

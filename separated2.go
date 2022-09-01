package spa

// Separated2 parses sep-separated sequence with at least two elements.
// See [Separated0] and [Separated1] for alternative parsers.
func Separated2(sep Parser, p Parser) Parser {
	return MinLen(2, Separated1(sep, p))
}

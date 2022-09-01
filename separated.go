package spa

// Separated0 parses sep-separated sequence with zero or more elements.
// See [Separated1] and [Separated2] for alternative parsers.
func Separated0(sep Parser, p Parser) Parser {
	return func(in string) Result {
		if r := p(in); r.Error == nil {
			return Tail(r.Value, Right(sep, p))(r.Input)
		} else {
			return Ok(in, make([]any, 0))
		}
	}
}

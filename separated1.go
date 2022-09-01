package spa

// Separated1 parses sep-separated sequence with at least one element.
// See [Separated0] and [Separated2] for alternative parsers.
func Separated1(sep Parser, p Parser) Parser {
	return func(in string) Result {
		if r := p(in); r.Error == nil {
			return Tail(r.Value, Right(sep, p))(r.Input)
		} else {
			return Errorf("expected at least 1 separated value; %v", r.Error)
		}
	}
}

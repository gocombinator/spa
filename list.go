package spa

// List matches sep-separated elements.
//
//	(p (sep p)*)? -> [..p]
func List(p, sep Parser) Parser {
	return func(in string) Result {
		if r := p(in); r.Err == nil {
			return valueTail(r.Value, Right(sep, p))(r.Input)
		} else {
			return Empty(in)
		}
	}
}

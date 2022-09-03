package spa

// Repeat matches at least once.
//
//	p+ -> [..p]
//
// See [Repeat0] for parser that matches 0 or more times.
func Repeat(p Parser) Parser {
	return func(in string) Result {
		if r := p(in); r.Err == nil {
			return valueTail(r.Value, p)(r.Input)
		} else {
			return r.Errorf("in repeat")
		}
	}
}

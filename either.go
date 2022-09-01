package spa

// Either matches a or b.
func Either(a, b Parser) Parser {
	return func(in string) Result {
		r := a(in)
		if r.Error == nil {
			return r
		}
		return b(in)
	}
}

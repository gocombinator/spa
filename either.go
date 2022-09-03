package spa

// Either tries first to match a, if it fails it matches b returning its result.
//
//	a / b -> a|b
func Either(a, b Parser) Parser {
	return func(in string) Result {
		r := a(in)
		if r.Err == nil {
			return r
		}
		return b(in)
	}
}

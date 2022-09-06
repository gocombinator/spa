package spa

// Either tries first to match a, if it fails it matches b returning its result.
//
//	a / b -> a|b
func Either(a, b Parser) Parser {
	return func(in string) (any, int, error) {
		if v, in_, err := a(in); err == nil {
			return v, in_, nil
		} else {
			return b(in)
		}
	}
}

package spa

// Until matches until provided parser succeeds or end of input is reached.
// Doesn't consume input of succeeding parser.
// Returns string with zero or more length.
//
//	(.*) (?p) -> $1
//
// Use [Min] to limit minimum length.
func Until(p Parser) Parser {
	return func(in string) (any, int, error) {
		var o = 0
		for {
			if _, w, err := p(in[o:]); err == nil && w > 0 {
				o += w
			} else {
				return in[:o], o, nil
			}
		}
		return Rest(in)
	}
}

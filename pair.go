package spa

// Pair `a / b -> [2][a,b]` matches pair of parsers in sequence returning 2-size array.
func Pair(a, b Parser) Parser {
	return func(in string) (any, int, error) {
		if v, w, err := a(in); err == nil {
			if v_, w_, err := b(in[w:]); err == nil {
				return []any{v, v_}, w + w_, nil
			} else {
				return nil, 0, err
			}
		} else {
			return nil, 0, err
		}
	}
}

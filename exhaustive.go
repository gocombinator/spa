package spa

// Exhaustive asserts that parser matches full input.
func Exhaustive(p Parser) Parser {
	return func(in string) Result {
		if r := p(in); r.Err == nil {
			if r.Input != "" {
				return r.Errorf("non exhaustive, %d chars left", len(r.Input))
			}
			return r
		} else {
			return r.Errorf("in exhaustive")
		}
	}
}

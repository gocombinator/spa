package spa

// Times matches parser n times.
// See [Repeat] and [Repeat0] for alternative parsers.
// See [MinLen] for related result constraint.
func Times(n int, p Parser) Parser {
	return func(in string) Result {
		var values []any
		for i := 0; i < n; i++ {
			if r := p(in); r.Error == nil {
				values = append(values, r.Value)
				in = r.Input
			} else {
				break
			}
		}
		if len(values) == n {
			return Ok(in, values)
		}
		return Errorf("expected %d times, got %d", n, len(values))
	}
}

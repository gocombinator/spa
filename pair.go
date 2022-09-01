package spa

// Pair matches pair of parsers in sequence returning 2-size array.
func Pair(p1, p2 Parser) Parser {
	return func(in string) Result {
		if r1 := p1(in); r1.Error == nil {
			if r2 := p2(r1.Input); r2.Error == nil {
				return Ok(r2.Input, [2]any{r1.Value, r2.Value})
			} else {
				return r2
			}
		} else {
			return r1
		}
	}
}

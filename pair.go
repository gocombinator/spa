package spa

// Pair `a / b -> [2][a,b]` matches pair of parsers in sequence returning 2-size array.
func Pair(a, b Parser) Parser {
	return func(in string) Result {
		if r1 := a(in); r1.Err == nil {
			if r2 := b(r1.Input); r2.Err == nil {
				return Ok(r2.Input, [2]any{r1.Value, r2.Value})
			} else {
				return r2
			}
		} else {
			return r1
		}
	}
}

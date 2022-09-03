package spa

// Times matches between n and m times.
//
//	p{n,m} -> [..p]
//
// See [Times], [Repeat] or [Repeat0] for alternative parsers.
// See [MinLen] for related result constraint.
func Times(n, m int, p Parser) Parser {
	return func(in string) (r Result) {
		var values []any
		r.Input = in
		for i := 0; i < n; i++ {
			if r = p(r.Input); r.Err == nil {
				values = append(values, r.Value)
			}
		}
		var l = len(values)
		if n <= l && l <= m {
			return Ok(in, values)
		}
		return r.Errorf("expected between %d-%d times, got %d last error", n, m, l)
	}
}

package spa

// NTimes matches n times.
//
//	p{n} -> []p
//
// See [Repeat] and [Repeat0] for alternative parsers.
// See [MinLen] for related result constraint.
func NTimes(n int, p Parser) Parser {
	return func(in string) Result {
		var err error
		var values []any
		for i := 0; i < n; i++ {
			if r := p(in); r.Err == nil {
				values = append(values, r.Value)
				in = r.Input
			} else {
				err = r.Err
				break
			}
		}
		if len(values) == n {
			return Ok(in, values)
		}
		return Errorf(in, "expected %d times, got %d; last error %w", n, len(values), err)
	}
}

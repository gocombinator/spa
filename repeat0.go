package spa

// Repeat0 returns parser that matches given parser zero or more times.
// See [Repeat] for parser that matches 1 or more times.
func Repeat0(p Parser) Parser {
	return func(in string) Result {
		var values []any
		for {
			if r := p(in); r.Error == nil {
				values = append(values, r.Value)
				in = r.Input
			} else {
				break
			}
		}
		return Ok(in, values)
	}
}

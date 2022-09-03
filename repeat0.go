package spa

// Repeat0 matches zero or more times.
//
//	p* -> [..p[nth]]
//
// See [Repeat] for parser that matches 1 or more times.
func Repeat0(p Parser) Parser {
	return func(in string) Result {
		var values = make([]any, 0)
		for {
			if r := p(in); r.Err == nil {
				values = append(values, r.Value)
				in = r.Input
			} else {
				break
			}
		}
		return Ok(in, values)
	}
}

package spa

// Seq matches parsers one by one in sequence.
func Seq(ps ...Parser) Parser {
	return func(in string) Result {
		var values []any
		for i, p := range ps {
			if r := p(in); r.Error == nil {
				values = append(values, r.Value)
				in = r.Input
			} else {
				return Errorf("expected %d/%d sequence; %v", i+1, len(ps), r.Error)
			}
		}
		return Ok(in, values)
	}
}

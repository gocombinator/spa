package spa

// Seq matches parsers one by one in sequence.
//
//	a / .. / z  -> [a, .., z]
func Seq(ps ...Parser) Parser {
	return func(in string) Result {
		var values []any
		for i, p := range ps {
			if r := p(in); r.Err == nil {
				values = append(values, r.Value)
				in = r.Input
			} else {
				return r.Errorf("in %d/%d seq", i+1, len(ps))
			}
		}
		return Ok(in, values)
	}
}

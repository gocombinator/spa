package spa

// Name decorates parser with name visible in parsing error stack.
func Name(name string, p Parser) Parser {
	return func(in string) Result {
		if r := p(in); r.Err == nil {
			return r
		} else {
			return r.Errorf("in %s", name)
		}
	}
}

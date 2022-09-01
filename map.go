package spa

// Map changes parser's result to new value.
func Map[T, U any](p Parser, f func(T) U) Parser {
	return func(in string) Result {
		if r := p(in); r.Error == nil {
			return Ok(r.Input, f(r.Value.(T)))
		} else {
			return r
		}
	}
}

package spa

// Make creates final parser.
func Make[T any](p Parser) func(string) (T, error) {
	return func(in string) (T, error) {
		if r := p(in); r.Error == nil {
			return r.Value.(T), nil
		} else {
			var zero T
			return zero, r.Error
		}
	}
}

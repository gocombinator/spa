package spa

func Then(f func(any, int, string) (any, int, error)) Mapper {
	return func(p Parser) Parser {
		return func(in string) (any, int, error) {
			if v, n, err := p(in); err == nil {
				return f(v, n, in)
			} else {
				return nil, 0, err
			}
		}
	}
}

package spa

func Catch(f func(error, string) (any, int, error)) Mapper {
	return func(p Parser) Parser {
		return func(in string) (any, int, error) {
			if v, n, err := p(in); err == nil {
				return v, n, nil
			} else {
				return f(err, in)
			}
		}
	}
}

package spa

import "fmt"

// As maps result to new value.
func As[T, U any](f func(T) U) Mapper {
	return func(p Parser) Parser {
		return func(in string) (any, int, error) {
			if v, n, err := p(in); err == nil {
				switch v := v.(type) {
				case nil:
					var zero T
					return f(zero), n, nil
				case T:
					return f(v), n, nil
				default:
					return nil, 0, fmt.Errorf("unexpected type %T (%v)", v, v)
				}
			} else {
				return nil, 0, err
			}
		}
	}
}

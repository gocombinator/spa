package spa

import "fmt"

// As maps result to new value.
func As[T, U any](p Parser, f func(T) U) Parser {
	return func(in string) Result {
		if r := p(in); r.Err == nil {
			switch v := r.Value.(type) {
			case nil:
				var zero T
				return Ok(r.Input, f(zero))
			case T:
				return Ok(r.Input, f(v))
			default:
				panic(fmt.Sprintf("unexpected type %T", v))
			}
		} else {
			return r
		}
	}
}

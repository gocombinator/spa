package spa

import "fmt"

// NTimes matches n times.
//
//	p{n} -> []p
//
// See [Repeat] and [Repeat0] for alternative parsers.
// See [MinLen] for related result constraint.
func NTimes(times int) func(Parser) Parser {
	return func(p Parser) Parser {
		return func(in string) (any, int, error) {
			var err error
			var values []any
			var o = 0
			for i := 0; i < times; i++ {
				if v, w, err := p(in); err == nil {
					values = append(values, v)
					o += w
				} else {
					break
				}
			}
			if len(values) == times {
				return values, o, nil
			}
			return nil, 0, fmt.Errorf("expected %d times, got %d; last error %w", times, len(values), err)
		}
	}
}

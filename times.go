package spa

import "fmt"

// Times matches between n and m times.
//
//	p{n,m} -> [..p]
//
// See [Times], [Repeat] or [Repeat0] for alternative parsers.
// See [MinLen] for related result constraint.
func Times(n, m int) Mapper {
	return func(p Parser) Parser {
		return func(in string) (any, int, error) {
			var values []any
			var o = 0
			for i := 0; i < m; i++ {
				if v, w, err := p(in[o:]); err == nil {
					values = append(values, v)
					o += w
				} else {
					break
				}
			}
			var l = len(values)
			if n <= l && l <= m {
				return values, o, nil
			}
			return nil, 0, fmt.Errorf("expected between %d-%d times, got %d", n, m, l)
		}
	}
}

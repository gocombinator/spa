package spa

import "fmt"

// Repeat matches at least once.
//
//	p+ -> [..p]
//
// See [Repeat0] for parser that matches 0 or more times.
func Repeat(p Parser) Parser {
	return func(in string) (any, int, error) {
		if v, w, err := p(in); err == nil {
			var values = make([]any, 1)
			values[0] = v
			var o = w
			for {
				if v, w, err := p(in[o:]); err == nil {
					values = append(values, v)
					o += w
				} else {
					break
				}
			}
			return values, o, nil
		} else {
			return nil, 0, fmt.Errorf("repeat, %w", err)
		}
	}
}

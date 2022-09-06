package spa

import "fmt"

// Seq matches parsers one by one in sequence.
//
//	a / .. / z  -> [a, .., z]
func Seq(ps ...Parser) Parser {
	return func(in string) (any, int, error) {
		var values []any
		var o = 0
		for i, p := range ps {
			if v, w, err := p(in[o:]); err == nil {
				values = append(values, v)
				o += w
			} else {
				return nil, 0, fmt.Errorf("in %d/%d seq, %w", i+1, len(ps), err)
			}
		}
		return values, o, nil
	}
}

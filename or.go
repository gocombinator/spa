package spa

import "fmt"

// Or `p[0] / .. / p[n-1] -> p[0]|..|p[n-1]` matches first parser from provided list.
func Or(ps ...Parser) Parser {
	return func(in string) (any, int, error) {
		for _, p := range ps {
			if v, w, err := p(in); err == nil {
				return v, w, nil
			}
		}
		return nil, 0, fmt.Errorf("in or none of %d parsers matched", len(ps))
	}
}

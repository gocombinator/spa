package spa

import "fmt"

// Exhaustive asserts that parser matches full input.
func Exhaustive(p Parser) Parser {
	return func(in string) (any, int, error) {
		if v, w, err := p(in); err != nil {
			return nil, 0, fmt.Errorf("in exhaustive, %w", err)
		} else if l := len(in); l != w {
			return nil, 0, fmt.Errorf("non exhaustive, %d/%d chars left", l-w, l)
		} else {
			return v, w, nil
		}
	}
}

package spa

import "fmt"

// Or matches first parser from provided list.
func Or(ps ...Parser) Parser {
	return func(in string) Result {
		for _, p := range ps {
			if r := p(in); r.Error == nil {
				return r
			}
		}
		return Errorf(fmt.Sprintf("none of %d parsers matched", len(ps)))
	}
}

package spa

// Repeat0 matches zero or more times.
//
//	p* -> [..p[nth]]
//
// See [Repeat] for parser that matches 1 or more times.
func Repeat0(p Parser) Parser {
	return func(in string) (any, int, error) {
		var values = make([]any, 0)
		var o = 0
		for {
			if v, w, err := p(in[o:]); err == nil {
				values = append(values, v)
				o += w
			} else {
				break
			}
		}
		return values, o, nil
	}
}

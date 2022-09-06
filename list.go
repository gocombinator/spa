package spa

// MaybeList matches sep-separated elements.
//
//	(p (sep p)*)? -> [..p]
func MaybeList(sep Parser) Mapper {
	return func(p Parser) Parser {
		return func(in string) (any, int, error) {
			if v, w, err := p(in); err == nil {
				var vs = make([]any, 1)
				vs[0] = v
				var o = w
				var q = Right(sep, p)
				for {
					if v, w, err := q(in[o:]); err == nil {
						vs = append(vs, v)
						o += w
					} else {
						break
					}
				}
				return vs, o, nil
			} else {
				return nil, 0, nil
			}
		}
	}
}

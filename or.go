package spa

// Or `p[0] / .. / p[n-1] -> p[0]|..|p[n-1]` matches first parser from provided list.
func Or(ps ...Parser) Parser {
	return func(in string) (r Result) {
		for _, p := range ps {
			if r = p(in); r.Err == nil {
				return r
			}
		}
		return r.Errorf("in or none of %d parsers matched (last error)", len(ps))
	}
}

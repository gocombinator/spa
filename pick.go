package spa

// Pick chooses single, ith result value or a list of result values if more indices are provided.
func Pick(p Parser, i int, indices ...int) Parser {
	return As(p, func(vs []any) any {
		if len(indices) == 0 {
			return vs[i]
		}
		var ps = make([]any, 1+len(indices))
		ps[0] = vs[i]
		for i, j := range indices {
			ps[1+i] = vs[j]
		}
		return ps
	})
}

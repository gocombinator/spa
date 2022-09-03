package spa

// Right matches pair of parsers in sequence returning second's result.
//
//	l / r -> r
//
// See [Left].
func Right(l Parser, r Parser) Parser {
	return As(Pair(l, r), func(vs [2]any) any {
		return vs[1]
	})
}

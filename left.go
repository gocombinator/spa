package spa

// Left matches pair of parsers in sequence returning first's result.
//
//	a / b -> a
func Left(a Parser, b Parser) Parser {
	return As(Pair(a, b), func(vs [2]any) any {
		return vs[0]
	})
}

package spa

// Left matches pair of parsers in sequence returning first one's result.
func Left(p Parser, q Parser) Parser {
	return Map(Pair(p, q), func(vs [2]any) any {
		return vs[0]
	})
}

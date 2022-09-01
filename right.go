package spa

// Right matches pair of parsers in sequence returning second one's result.
func Right(p Parser, q Parser) Parser {
	return Map(Pair(p, q), func(vs [2]any) any {
		return vs[1]
	})
}

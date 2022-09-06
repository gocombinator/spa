package spa

// Right matches pair of parsers in sequence returning second's result.
//
//	l / r -> r
//
// See [Left].
func Right(l, r Parser) Parser {
	return Pipe(
		Pair(l, r),
		Pick(1),
	)
}

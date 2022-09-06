package spa

// Wrap matches parser sorrounded by left and right.
//
//	l p r -> p
func Wrap(l, r Parser) Mapper {
	return func(p Parser) Parser {
		return Pipe(
			Seq(l, p, r),
			Pick(1),
		)
	}
}

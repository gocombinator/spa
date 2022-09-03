package spa

// Wrap matches parser sorrounded by left and right.
//
//	l p r -> p
func Wrap(p, l, r Parser) Parser {
	return Pick(Seq(l, p, r), 1)
}

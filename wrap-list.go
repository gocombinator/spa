package spa

// WrapList matches wrapped list.
//
//	l (p (sep p)*)? r -> [..p]
func WrapList(sep, l, r Parser) Mapper {
	return func(p Parser) Parser {
		return Pipe(
			p,
			MaybeList(sep),
			Wrap(l, r),
		)
	}
}

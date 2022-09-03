package spa

// WrapList matches wrapped list.
//
//	l (p (sep p)*)? r -> [..p]
func WrapList(p, sep, l, r Parser) Parser {
	return Wrap(List(p, sep), l, r)
}

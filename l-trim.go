package spa

// LTrim ignores leading part.
// See [Trim].
func LTrim(ps ...Parser) Mapper {
	if len(ps) == 0 {
		return Wrap(Ws0, Empty)
	}
	if len(ps) == 1 {
		return Wrap(ps[0], Empty)
	}
	var p = Repeat0(Or(ps...))
	return Wrap(p, Empty)
}

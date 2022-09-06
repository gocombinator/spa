package spa

// Trim ignores leading and trailing part.
//
// If there are no trim parser(s) provided, [Ws0] is used.
//
// If single trim parser is provied, it is matched at both ends as is.
// For optional trim you want to pass single zero-width-succeeding parser.
//
// If more trim parsers are provided, they're [Or]'ed and [Repeat0]'ed.
// In this case none of trim parsers should be zero-width-succeeding.
func Trim(ps ...Parser) Mapper {
	if len(ps) == 0 {
		return Wrap(Ws0, Ws0)
	}
	if len(ps) == 1 {
		return Wrap(ps[0], ps[0])
	}
	var p = Repeat0(Or(ps...))
	return Wrap(p, p)
}

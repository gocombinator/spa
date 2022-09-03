package spa

// Trim ignores leading and trailing part.
// If no trim parsers are provided, whitespaces are trimmed.
func Trim(p Parser, trims ...Parser) Parser {
	return func(in string) Result {
		if len(trims) == 0 {
			return Pick(Seq(Ws0, p, Ws0), 1)(in)
		}
		var q = Repeat0(Or(trims...))
		return Pick(Seq(q, p, q), 1)(in)
	}
}

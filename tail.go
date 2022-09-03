package spa

// Tail matches head with zero or more tail repetitions.
//
//	h t* -> [h, ..t]
func Tail(h, t Parser) Parser {
	return Seq(h, Repeat0(t))
}

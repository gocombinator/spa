package spa

// Repeat returns parser that matches given parser at least once.
func Repeat(p Parser) Parser {
	return MinLen(1, Repeat0(p))
}

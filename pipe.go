package spa

// Pipe matches a sequence of parsers.
func Pipe(p Parser, fs ...func(Parser) Parser) Parser {
	for _, f := range fs {
		p = f(p)
	}
	return p
}

package spa

// Nth maps parser's result to nth slice value.
func Nth(i int, p Parser) Parser {
	return Map(p, func(r any) any { return r.([]any)[i] })
}

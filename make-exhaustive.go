package spa

// MakeExhaustive creates final, exhaustive parser.
func MakeExhaustive[T any](p Parser) func(string) (T, error) {
	return Make[T](Exhaustive(p))
}

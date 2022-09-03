package spa

// Until matches until provided parser succeeds or end of input is reached.
// Doesn't consume input of succeeding parser.
// Returns string with zero or more length.
//
//	(.*) (?p) -> $1
//
// Use [Min] to limit minimum length.
func Until(p Parser) Parser {
	return func(in string) Result {
		for i := range in {
			if p(in[i:]).Err == nil {
				return Eat(in, i)
			}
		}
		return Rest(in)
	}
}

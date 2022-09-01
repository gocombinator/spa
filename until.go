package spa

// Until matches string until provided parser succeeds (successful match is not included)
// or end of string is reached (returns success).
func Until(p Parser) Parser {
	return func(in string) Result {
		for i := range in {
			if p(in[i:]).Error == nil {
				return Ok(in[i:], in[:i])
			}
		}
		return Ok("", in)
	}
}

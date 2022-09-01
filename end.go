package spa

// End matches end of input.
func End() Parser {
	return func(in string) Result {
		if in == "" {
			return Ok("", nil)
		} else {
			return Errorf("expected end of input")
		}
	}
}

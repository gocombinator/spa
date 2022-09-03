package spa

// Rest matches everything until end of input.
func Rest(in string) Result {
	return Ok("", in)
}

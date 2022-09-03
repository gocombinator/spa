package spa

// Eat consumes n bytes.
func Eat(in string, n int) Result {
	return Ok(in[n:], in[:n])
}

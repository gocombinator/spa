package spa

// Empty returns successful nil result without consuming input.
func Empty(in string) Result {
	return Ok(in, nil)
}

package spa

// End matches end of input.
func End(in string) Result {
	if in == "" {
		return Empty("")
	} else {
		return Errorf(in, "expected end of input")
	}
}

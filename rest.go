package spa

// Rest matches everything until end of input.
func Rest(in string) (any, int, error) {
	return in, len(in), nil
}

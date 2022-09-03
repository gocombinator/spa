package spa

// Make creates final parse function matching input exhaustively.
// See [Trim] to make parse function more loose ignoring trailing and leading part.
func Make[T any](p Parser) func(string) (T, error) {
	return func(in string) (T, error) {
		if r := Exhaustive(p)(in); r.Err == nil {
			return r.Value.(T), nil
		} else {
			// fmt.Fprintf(os.Stderr, "--<trace>\n%s\n--</trace>\n", r.Trace(in))
			var zero T
			return zero, r.Err
		}
	}
}

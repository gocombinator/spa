package spa

import "fmt"

// Make creates final parse function matching input exhaustively.
// See [Trim] to make parse function more loose ignoring trailing and leading part.
func Make[T any](p Parser) func(string) (T, error) {
	return func(in string) (T, error) {
		if v, _, err := Exhaustive(p)(in); err == nil {
			switch v := v.(type) {
			case nil:
				var zero T
				return zero, nil
			case T:
				return v, nil
			default:
				var zero T
				return zero, fmt.Errorf("unexpected type %T", v)
			}
		} else {
			// fmt.Fprintf(os.Stderr, "--<trace>\n%s\n--</trace>\n", r.Trace(in))
			var zero T
			return zero, err
		}
	}
}

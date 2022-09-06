package spa

import "fmt"

// End matches end of input.
func End(in string) (any, int, error) {
	if in == "" {
		return nil, 0, nil
	} else {
		return nil, 0, fmt.Errorf("expected end of input")
	}
}

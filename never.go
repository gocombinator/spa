package spa

import "fmt"

// Never always fails.
func Never(in string) (any, int, error) {
	return nil, 0, fmt.Errorf("never")
}

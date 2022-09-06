package spa

import "fmt"

// Name decorates error message with name.
func Name(name string) Mapper {
	return Catch(func(err error, _ string) (any, int, error) {
		return nil, 0, fmt.Errorf("%s: %w", name, err)
	})
}

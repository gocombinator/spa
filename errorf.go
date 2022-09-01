package spa

import "fmt"

// Errorf returns error result.
func Errorf(format string, args ...any) Result {
	return Result{Error: fmt.Errorf(format, args...)}
}

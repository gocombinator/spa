package spa

import "fmt"

// Errorf returns error result.
func Errorf(in string, format string, args ...any) Result {
	return Result{Input: in, Err: fmt.Errorf(format, args...)}
}

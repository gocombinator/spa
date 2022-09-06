package spa

type Tracer interface {
	Trace() string
}

// Result represents [Parser] result.
type Result struct {
	Input string
	Value any
}

// func (r Result) Errorf(format string, args ...any) (any, int, error) {
// 	args = append(args, r)
// 	return Result{Input: r.Input, Err: fmt.Errorf(format+"; %w", args...)}
// }

// // Trace implements [Tracer] interface.
// func (r Result) Trace(in string) string {
// 	var b strings.Builder
// 	fmt.Fprintf(&b, "%s\n", in)
// 	var err = r.Err
// 	var n = len(in)
// 	for {
// 		var m = len(r.Input)
// 		var i = n - m
// 		var margin = strings.Repeat(" ", i)
// 		fmt.Fprintf(&b, "%s^\n", margin)
// 		fmt.Fprintf(&b, "%s%v\n", margin, err)
// 		if err = errors.Unwrap(err); err == nil {
// 			break
// 		}
// 	}
// 	return b.String()
// }

package spa

// Tail appends 0 or more repeated parser results to optional first value.
// First value is not included if it's nil.
func Tail(firstValue any, p Parser) Parser {
	return func(in string) Result {
		var values = make([]any, 0)
		if firstValue != nil {
			values = append(values, firstValue)
		}
		if r := Repeat0(p)(in); r.Error == nil {
			values = append(values, r.Value.([]any)...)
			return Ok(r.Input, values)
		} else {
			return r
		}
	}
}

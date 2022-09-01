package spa

// Ok returns successful result.
func Ok(input string, value any) Result {
	return Result{Input: input, Value: value}
}

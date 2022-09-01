package internal

import "fmt"

// Len returns length of any value.
func Len(value any) int {
	switch v := value.(type) {
	case string:
		return len(v)
	case []any:
		return len(v)
	default:
		panic(fmt.Sprintf("unsupported type: %T", value))
	}
}

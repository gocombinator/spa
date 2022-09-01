package internal

import (
	"fmt"
	"strings"
)

func Join(values any, sep string) string {
	switch values := values.(type) {
	case []string:
		return strings.Join(values, sep)
	case []any:
		return Join(Slicemap(values, func(v any) string {
			switch v := v.(type) {
			case string:
				return v
			case fmt.Stringer:
				return v.String()
			default:
				panic(fmt.Sprintf("unexpected type %T", v))
			}
		}), sep)
	default:
		panic(fmt.Sprintf("unsupported concat type %T", values))
	}
}

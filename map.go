package spa

import "fmt"

// Map changes result from slice of key values to map.
//
//	p -> [..[k,v]] -> map[k]v
func Map[K comparable, V any]() Mapper {
	return As(func(kvs []any) map[K]V {
		var m = make(map[K]V, len(kvs))
		for _, kv := range kvs {
			switch kv := kv.(type) {
			case nil:
			case []any:
				m[kv[0].(K)] = kv[1].(V)
			case [2]any:
				m[kv[0].(K)] = kv[1].(V)
			default:
				panic(fmt.Sprintf("unexpected type %T", kv))
			}
		}
		return m
	})
}

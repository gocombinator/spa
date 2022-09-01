package spa

// Label converts slice result into a map.
func Label(p Parser, labels ...string) Parser {
	return Map(p, func(v []any) map[string]any {
		var m = make(map[string]any)
		for i, label := range labels {
			if label != "" {
				m[label] = v[i]
			}
		}
		return m
	})
}

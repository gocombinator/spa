package internal

import (
	"fmt"
	"strings"
)

func ReadablePatterns(patterns ...string) string {
	var b strings.Builder
	for i, pattern := range patterns {
		pattern := Slicemap([]rune(pattern), func(r rune) string {
			if 32 < r && r < 126 && r != '/' {
				return string(r)
			} else {
				if r <= 0xff {
					return fmt.Sprintf("%2x", r)
				} else if r <= 0xffff {
					return fmt.Sprintf("%4x", r)
				} else {
					return fmt.Sprintf("%8x", r)
				}
			}
		})
		if i > 0 {
			if i%2 == 1 {
				b.WriteString("-")
			} else {
				b.WriteString(" / ")
			}
		}
		b.WriteString(strings.Join(pattern, " / "))
	}
	return b.String()
}

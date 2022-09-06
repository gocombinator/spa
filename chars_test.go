package spa_test

import (
	"github.com/gocombinator/spa"
)

func ExampleChars() {
	var s = spa.Chars("a", "z")
	var n = spa.Chars("0", "9")
	var p = spa.Repeat(spa.Either(s, n))
	print(p, "abc123")
	// Output:
	// abc123 -> [abc 123]
}

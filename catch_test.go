package spa_test

import (
	"fmt"

	"github.com/gocombinator/spa"
)

func ExampleCatch() {
	var id = spa.Chars("a", "z")
	var ids = spa.Pipe(
		id,
		spa.Catch(func(err error, in string) (any, int, error) {
			if in == "" {
				return nil, 0, err
			}
			return spa.Until(id)(in)
		}),
	)
	var p = spa.Repeat(ids)
	var v, _, _ = p("abc-123:def456")
	fmt.Println(v)
	// Output:
	// [abc -123: def 456]
}

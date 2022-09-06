package spa_test

import (
	"fmt"
	"strconv"

	"github.com/gocombinator/spa"
)

func ExampleAs() {
	var num = spa.Pipe(
		spa.Chars("0", "9"),
		spa.As(func(s string) int {
			n, _ := strconv.Atoi(s)
			return n
		}),
	)
	var parse = spa.Make[int](num)
	if r, err := parse("123"); err == nil {
		fmt.Printf("%d (%T)", r, r)
		// Output:
		// 123 (int)
	} else {
		fmt.Printf("%s", err)
	}
}

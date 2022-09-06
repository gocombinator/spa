package spa_test

import (
	"fmt"

	"github.com/gocombinator/spa"
)

func ExamplePratt() {
	var primary = spa.Chars("0", "z")
	var infix = spa.Or(
		spa.Pipe(spa.String("+", "-"), spa.As(func(s string) spa.PrecedenceResult {
			return spa.PrecedenceResult{Value: s, Left: 1, Right: 2}
		})),
		spa.Pipe(spa.String("*", "/"), spa.As(func(s string) spa.PrecedenceResult {
			return spa.PrecedenceResult{Value: s, Left: 3, Right: 4}
		})),
	)
	var prefix = spa.Never
	var postfix = spa.Never
	var expr = spa.Pratt(primary, prefix, infix, postfix)
	var parser = spa.Exhaustive(expr)
	if r, _, err := parser(`1+2*3`); err == nil {
		fmt.Println(r)
	} else {
		fmt.Println("error:", err)
	}
	if r, _, err := parser(`1*2+3`); err == nil {
		fmt.Println(r)
	} else {
		fmt.Println("error:", err)
	}
	// Output:
	// (+ 1 (* 2 3))
	// (+ (* 1 2) 3)
}

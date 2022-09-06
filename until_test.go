package spa_test

import (
	"fmt"

	"github.com/gocombinator/spa"
)

func ExampleUntil() {
	var beg = spa.Char("<")
	var end = spa.Char(">")
	var tag = spa.Pipe(
		spa.Seq(
			beg,
			spa.Until(end),
			end,
		),
		spa.Pick(1),
		spa.As(func(s string) string {
			return fmt.Sprintf("(%s)", s)
		}),
	)
	var tags = spa.Repeat(tag)
	var parse = spa.Make[any](tags)
	if v, err := parse("<a><b><><cde>"); err == nil {
		fmt.Println(v)
	} else {
		fmt.Println(err)
	}
	// Output:
	// [(a) (b) () (cde)]
}

func ExampleUntil_tillEnd() {
	var p = spa.Until(spa.Char("!"))
	if v, _, err := p("abc"); err == nil {
		fmt.Println(v)
	} else {
		fmt.Println(err)
	}
	// Output:
	// abc
}

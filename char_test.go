package spa_test

import (
	"fmt"

	"github.com/gocombinator/spa"
)

func ExampleChar() {
	var c = spa.Char(
		"ac", "e",
		"gh", "j",
	)
	var p = func(in string) {
		if v, _, err := c(in); err == nil {
			fmt.Printf("%s -> %v\n", in, v)
		} else {
			fmt.Printf("%s -> %v\n", in, err)
		}
	}
	for i := 'a'; i <= 'k'; i++ {
		p(string(i))
	}
	// Output:
	// a -> a
	// b -> expected one of a / c-e / g / h-j chars
	// c -> c
	// d -> d
	// e -> e
	// f -> expected one of a / c-e / g / h-j chars
	// g -> g
	// h -> h
	// i -> i
	// j -> j
	// k -> expected one of a / c-e / g / h-j chars
}

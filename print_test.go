package spa_test

import (
	"fmt"

	"github.com/gocombinator/spa"
)

// print prints parser's output for provided input string.
func print(p spa.Parser, in string) {
	if v, _, err := p(in); err == nil {
		fmt.Printf("%s -> %v\n", in, v)
	} else {
		fmt.Printf("%s -> %v\n", in, err)
	}
}

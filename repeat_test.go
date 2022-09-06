package spa_test

import (
	"fmt"

	"github.com/gocombinator/spa"
)

func ExampleRepeat() {
	if v, w, err := spa.Repeat(spa.Char("a", "z"))("abc123"); err == nil {
		fmt.Println(v, w)
	} else {
		fmt.Println(err)
	}
	if v, w, err := spa.Repeat(spa.Char("a", "z"))("123"); err == nil {
		fmt.Println(v, w)
	} else {
		fmt.Println(err, w)
	}
	// Output:
	// [a b c] 3
	// repeat, expected one of a-z chars 0
}

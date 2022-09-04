package spa_test

import (
	"fmt"

	"github.com/gocombinator/spa"
)

func ExampleMap() {
	var id = spa.Chars("a", "z")
	var kv = spa.Pick(spa.Seq(id, spa.Char("="), id), 0, 2)
	var kvs = spa.List(kv, spa.Ws1)
	var parser = spa.Map[string, string](kvs)
	var parse = spa.Make[map[string]string](parser)
	if r, err := parse("a=b c=d"); err == nil {
		fmt.Println(r)
		// Output: map[a:b c:d]
	} else {
		fmt.Println(err)
	}
}

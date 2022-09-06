package spa_test

import (
	"fmt"

	"github.com/gocombinator/spa"
)

func ExampleMap() {
	var id = spa.Chars("a", "z")
	var kv = spa.Pipe(
		spa.Seq(id, spa.Char("="), id),
		spa.Pick(0, 2),
	)
	var kvs = spa.Pipe(
		kv,
		spa.MaybeList(spa.Ws1),
	)
	var parser = spa.Pipe(
		kvs,
		spa.Map[string, string](),
	)
	var parse = spa.Make[map[string]string](parser)
	if r, err := parse("a=b c=d"); err == nil {
		fmt.Println(r)
		// Output: map[a:b c:d]
	} else {
		fmt.Println(err)
	}
}

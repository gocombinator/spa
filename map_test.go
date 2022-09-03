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
	var final = spa.Exhaustive(parser)
	var parse = spa.Make[map[string]string](final)
	if r, err := parse("a=b c=d"); err == nil {
		fmt.Println(r)
		// Output: map[a:b c:d]
	} else {
		fmt.Println(err)
	}
}

// func ExampleMap_xml() {

// 	// Construct parser for subset of xml.

// 	// Name of a tag.
// 	var name = spa.Chars("a", "z")

// 	// Attribute of a tag.
// 	var attr = spa.Pick(spa.Seq(name, spa.Char("="), json.String), 0, 2)

// 	// Assoc is used here to map attributes into single mapping.
// 	var attrs = spa.Map[string, any](spa.List(attr, spa.Ws1))

// 	var tag = spa.Wrap(
// 		spa.Label(spa.Seq(name, spa.Ws1, attrs), "name", "", "attrs"),
// 		spa.Char("<"),
// 		spa.String("/>"),
// 	)
// 	var tags = spa.Repeat(tag)
// 	var parse = spa.Make[any](spa.Exhaustive(tags))

// 	// Parse example file.
// 	if r, err := parse(`<xyz foo="a" bar="b"/>`); err == nil {
// 		fmt.Println(r)
// 		// Output: [map[attrs:map[bar:"b" foo:"a"] name:xyz]]
// 	} else {
// 		panic(err)
// 	}
// }

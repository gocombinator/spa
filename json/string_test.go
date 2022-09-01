package json_test

import (
	"testing"

	"github.com/gocombinator/spa"
	"github.com/gocombinator/spa/json"
)

func TestString(t *testing.T) {

	var same = func(p spa.Parser, inout string) {
		var p_ = spa.Make[string](spa.Exhaustive(p))
		if r, err := p_(inout); err == nil {
			if r != inout {
				t.Errorf(`expected %s, got %s`, inout, r)
			}
		} else {
			t.Error(err)
		}
	}

	same(json.String, `""`)
	same(json.String, `"hello"`)

}

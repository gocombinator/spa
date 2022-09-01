package spa_test

import (
	"testing"

	"github.com/gocombinator/spa"
)

func TestChar(t *testing.T) {
	p := spa.Make[string](spa.Exhaustive(spa.Char("\"")))
	if r, err := p("\""); err == nil {
		if r != "\"" {
			t.Errorf("expected \", got %s", r)
		}
	} else {
		t.Error(err)
	}
}

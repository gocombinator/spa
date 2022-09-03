package internal_test

import (
	"testing"

	"github.com/gocombinator/spa/internal"
)

func TestReadablePatterns(t *testing.T) {
	if internal.ReadablePatterns("0", "f", "A", "F") != "0-f / A-F" {
		t.Fail()
	}
}

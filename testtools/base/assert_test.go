package base

import (
	"strings"
	"testing"
)

// A poor assertion function.
func assertEqual(t *testing.T, x, y interface{}) {
	t.Helper()
	if x != y {
		t.Errorf("Not Equal, %d %d", x, y)
	}
}

func TestSplit(t *testing.T) {
	words := strings.Split("a:b:c:d", ":")
	assertEqual(t, len(words), 4)
}

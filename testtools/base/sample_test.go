package base

import "testing"

func TestSample(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
}

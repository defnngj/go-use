package dataconversion

import (
	"testing"
)

func TestStringToIntSlice(t *testing.T) {
	str := "1,2,3"
	ret, err := StringToIntSlice(str, ",")
	if err != nil {
		t.Errorf("string: '%s', split:'%s'", str, ",")
	}
	t.Logf("%s to -> %d", str, ret)
}

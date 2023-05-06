package dataconversion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringToIntSlice(t *testing.T) {
	str := "1,2,3"
	ret, err := StringToIntSlice(str, ",")
	if err != nil {
		t.Errorf("string: '%s', split:'%s'", str, ",")
	}
	t.Logf("%s to -> %d", str, ret)
}

func TestInIntList(t *testing.T) {
	intList := []int{1, 2, 3, 4, 5}
	i := 3
	ret := InIntList(i, intList)
	t.Logf("%d to -> %d", i, intList)
	assert.Equal(t, ret, true)
}

func TestInStrList(t *testing.T) {
	intList := []string{"apple", "banana", "peach", "watermelon", "orange"}
	s := "peach"
	ret := InStrList(s, intList)
	t.Logf("%s to -> %s", s, intList)
	assert.Equal(t, ret, true)
}

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

func TestInt64SliceToStringSlice(t *testing.T) {
	var intList []int64
	intList = append(intList, 1)
	intList = append(intList, 2)
	intList = append(intList, 3)
	ret := Int64SliceToStringSlice(intList)
	t.Logf("%d to -> %s", intList, ret)
}

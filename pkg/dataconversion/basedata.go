package dataconversion

import (
	"strconv"
	"strings"
)

func StringToIntSlice(str string, sep string) ([]int, error) {
	if len(str) == 0 {
		return nil, nil
	}
	l := strings.Split(str, sep)
	res := make([]int, 0, len(l))
	for _, v := range l {
		i, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		res = append(res, i)
	}

	return res, nil
}

func Int64SliceToStringSlice(source []int64) []string {
	var res []string
	for _, s := range source {
		res = append(res, strconv.FormatInt(s, 10))
	}
	return res
}

func InIntList(numberToFind int, integerList []int) bool {
	for _, number := range integerList {
		if number == numberToFind {
			return true
		}
	}
	return false
}

func InStrList(strToFind string, strList []string) bool {
	for _, number := range strList {
		if number == strToFind {
			return true
		}
	}
	return false
}

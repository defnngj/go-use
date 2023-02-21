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

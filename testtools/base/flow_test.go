package base

import (
	"testing"
)

type Test struct {
	in  int
	out string
}

var testData = []Test{
	{55, "不及格"},
	{61, "及格"},
	{78, "良好"},
	{91, "优秀"},
}

func TestStudentMark(t *testing.T) {
	for i, test := range testData {
		mark := StudentMark(test.in)
		if mark != test.out {
			t.Errorf("%d: StudentMark(%d)=%s; want %s", i, test.in, mark, test.out)
		}
	}
}

package base

import (
	"errors"
	"fmt"
)

// StudentMark 计算学生成绩
func StudentMark(num int) string {
	fmt.Println("你的分数是:", num)
	var ret string
	if num > 90 {
		ret = "优秀"
	} else if num > 70 {
		ret = "良好"
	} else if num > 60 {
		ret = "及格"
	} else {
		ret = "不及格"
	}
	fmt.Println("你的评级为:", ret)
	return ret
}

// Addition 两数做加法运算
func Addition(num1 int, num2 int) int {
	fmt.Println("Addition: ", num1, "+", num2)
	var sum int = num1 + num2
	return sum
}

// Division 实现除法运算
func Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不合法，不能为0")
	}

	return a / b, nil
}

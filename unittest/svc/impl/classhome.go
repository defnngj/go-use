package impl

import "fmt"

type ClassHomelnf interface {
	LeanC(string) (string, error)
	LeanCPlus(string) (string, error)
	LeanGo(string) (string, error)
}

var _ClassHomelnf = (*ClassHome)(nil)

type ClassHome struct {
}

func NewClassHome() *ClassHome {
	return &ClassHome{}
}

func (c ClassHome) LeanC(s string) (string, error) {
	fmt.Println(s, "在课堂学c")
	return "classHome-C", nil
}

func (c ClassHome) LeanCPlus(s string) (string, error) {
	fmt.Println(s, "在课堂学C++")
	return "classHome-C++", nil
}

func (c ClassHome) LeanGo(s string) (string, error) {
	fmt.Println(s, "在课堂学GO")
	return "classHome-Go", nil
}

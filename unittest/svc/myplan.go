package svc

import (
	"fmt"
	"go-use/unittest/inf"
)

type MyLeanProxy interface {
	LeanGo(string) (string, error)
}

type MyPlan struct {
	LeanEvent MyLeanProxy
}

func NewMyPlanlmpl(lean inf.Lean) MyPlan {
	return MyPlan{LeanEvent: lean}
}

func (p MyPlan) Lean(req string) {
	c, err := p.LeanEvent.LeanGo(req)
	if err != nil {
		fmt.Println(err)
		fmt.Println(c)
	}

}

package main

import (
	"go-use/unittest/svc"
	"go-use/unittest/svc/impl"
)

func main() {
	svc.NewMyPlanlmpl(impl.NewClassHome()).Lean("虫师")
	svc.NewMyPlanlmpl(impl.NewBiliBili()).Lean("虫师")
}

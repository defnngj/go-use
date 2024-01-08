package impl

import "fmt"

type BiliBililnf interface {
	LeanC(string) (string, error)
	LeanGo(string) (string, error)
	LeanCPlus(string) (string, error)
}

var _BiliBililnf = (*BiliBili)(nil)

type BiliBili struct {
}

func NewBiliBili() *BiliBili {
	return &BiliBili{}
}

func (b BiliBili) LeanC(s string) (string, error) {
	fmt.Println(s, "在B站学c")
	return "BiliBili-c", nil
}

func (b BiliBili) LeanGo(s string) (string, error) {
	fmt.Println(s, "在B站学Go")
	return "BiliBili-go", nil
}

func (b BiliBili) LeanCPlus(s string) (string, error) {
	fmt.Println(s, "在B站学C++")
	return "BiliBili-c", nil
}

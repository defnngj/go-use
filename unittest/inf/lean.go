package inf

type Lean interface {
	LeanC(string) (string, error)
	LeanCPlus(string) (string, error)
	LeanGo(string) (string, error)
}

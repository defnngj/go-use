package pathconversion

import (
	"testing"
)

func TestStringToIntSlice(t *testing.T) {
	exePath := GetExePath()
	absPath := GetAbsPath()
	workingDirPath := GetWorkingDirPath()

	t.Logf("exePath -> %s", exePath)
	t.Logf("absPath -> %s", absPath)
	t.Logf("orkingDirPath to -> %s", workingDirPath)
}

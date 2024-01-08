package pathconversion

import (
	"os"
	"path/filepath"
)

func GetExePath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exePath := filepath.Dir(ex)

	return exePath
}

func GetAbsPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	return dir
}

func GetWorkingDirPath() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return dir
}

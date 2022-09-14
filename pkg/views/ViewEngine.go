package views

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GetTemplatePath() (s string, err error) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	wd = filepath.Dir(wd)
	wd = strings.Replace(wd, "cmd", "pkg/views/", -1)

	return wd, err

}

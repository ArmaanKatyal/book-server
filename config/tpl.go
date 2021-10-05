package config

import (
	"html/template"
	"os"
)

var TPL *template.Template

func init() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	TPL = template.Must(template.ParseGlob(wd + "/mongoWebApp/templates/*"))
}

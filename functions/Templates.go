package functions

import (
	"bytes"
	"html/template"
	"io"
	"log"
)

var templates *template.Template

func InitTemplates() {
	var err error
	templates, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatal(err)
	}
}

func Exec(name string, wr io.Writer, data any) error {
	var buffer bytes.Buffer

	err := templates.ExecuteTemplate(&buffer, name, data)
	if err != nil {
		return err
	}

	buffer.WriteTo(wr)

	return nil
}

package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", " Khoi")
	if err != nil {
		log.Fatalln(err)
	}
}

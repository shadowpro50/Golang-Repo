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
	name := []string{"Johnny", "Cassie", "Scorpion"}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", name)
	if err != nil {
		log.Fatalln(err)
	}
}

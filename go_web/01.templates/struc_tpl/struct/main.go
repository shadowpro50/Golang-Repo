package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type Person struct {
	Name   string
	Age    int
	Gender string
}

func init() {
	tpl = template.Must(template.ParseGlob("*"))
}

func main() {
	person1 := Person{"Hanzo", 23, "male"}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", person1)
	if err != nil {
		log.Fatalln(err)
	}
}

package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

type Page struct {
	Title   string
	Heading string
	Input   string
}

func init() {
	tpl = template.Must(template.ParseGlob("*"))
}

func main() {
	home := Page{
		Title:   "Hello, world",
		Heading: "Nothing is escaped",
		Input:   "<script>alert</script>",
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", home)
	if err != nil {
		log.Fatalln(err)
	}
}

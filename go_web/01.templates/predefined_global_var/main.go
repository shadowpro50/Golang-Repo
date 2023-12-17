package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	xs := []string{"one", "two", "three"}

	err := tpl.Execute(os.Stdout, xs)
	if err != nil {
		log.Fatalln(err)
	}
}

package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type person struct {
	Name string
	Age  int
}

func (p person) AgeDbl() int {
	return p.Age * 2
}
func (p person) AgeByFour(x int) int {
	return x * 2
}
func init() {
	tpl = template.Must(template.ParseGlob("*"))
}

func main() {
	p := person{
		Name: "shadowpro50",
		Age:  18,
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", p)
	if err != nil {
		log.Fatalln(err)
	}
}

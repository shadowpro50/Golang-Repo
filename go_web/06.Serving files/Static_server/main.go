package main

import (
	"html/template"
	"net/http"
)

// func main() {
// 	http.ListenAndServe(":8080", http.FileServer(http.Dir("web")))
// }

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("web/*"))
}

func index(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "index.html", nil)
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"html/template"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request)  {
	data := "Zam1r"
	tpl.ExecuteTemplate(w, "index.gohtml", data)
}
var tpl *template.Template
func init()  {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}
func main() {

	http.HandleFunc("/", root)

	http.ListenAndServe(":8080", nil)
}

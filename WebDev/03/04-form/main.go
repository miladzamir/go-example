package main

import (
	"html/template"
	"log"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	tpl.ExecuteTemplate(w, "page.gohtml", r.Form)
}

var tpl *template.Template
func init()  {
	tpl = template.Must(template.ParseFiles("page.gohtml"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}

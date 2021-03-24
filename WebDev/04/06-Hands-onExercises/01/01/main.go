package main

import (
	"io"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w, "ROOT")
}
func dog(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w, "DOG")
}
func me(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w, "MILAD")
}
func main() {

	http.HandleFunc("/", root)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}

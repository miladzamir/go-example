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

	http.Handle("/", http.HandlerFunc(root))
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/me/", http.HandlerFunc(me))

	http.ListenAndServe(":8080", nil)
}

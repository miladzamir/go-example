package main

import (
	"io"
	"net/http"
)

func d(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "DOG")
}

func c(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "CAT")
}
func main() {

	http.HandleFunc("/dog/", d)
	http.HandleFunc("/cat", c)

	http.ListenAndServe(":8080", nil)
}

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", root)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
func root(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.URL)
	fmt.Fprintln(w, "SEE LOG IN TERMINAL")
}

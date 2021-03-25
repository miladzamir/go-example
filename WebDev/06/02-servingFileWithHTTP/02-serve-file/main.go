package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/image", image)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
	<img src="/image" />
	`)
}
func image(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "Image_created_with_a_mobile_phone.png")
}

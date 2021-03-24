package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)

	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
	<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/b/b6/Image_created_with_a_mobile_phone.png/800px-Image_created_with_a_mobile_phone.png" />
	`)
}
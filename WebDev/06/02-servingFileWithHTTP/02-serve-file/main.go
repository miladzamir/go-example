package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/image", image)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
	<img src="/image" />
	`)
}
func image(w http.ResponseWriter, r *http.Request)  {
	f, err := os.Open("Image_created_with_a_mobile_phone.png")
	if err != nil {
		http.Error(w, "File Not Found", 404)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "File Not Found", 404)
	}
	http.ServeContent(w,r,fi.Name(),fi.ModTime(),f)
}
package main

import (
	"io"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func main() {
	http.HandleFunc("/", root)
	http.ListenAndServe(":3000", nil)
}

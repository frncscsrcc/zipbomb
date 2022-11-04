package main

import (
	"io"
	"net/http"
	"os"
)

func bomb(w http.ResponseWriter, req *http.Request) {
	dat, _ := os.Open("data")
	w.Header().Set("Content-Encoding", "gzip")
	io.Copy(w, dat)
}

func main() {
	http.HandleFunc("/", bomb)
	http.ListenAndServe(":8000", nil)
}

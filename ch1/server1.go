package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // each request calls header
	http.ListenAndServe("localhost:8000", nil)
}

// Handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

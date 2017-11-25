package main

import (
	"fmt"
	"net/http"
)

var body = "Nothing to see here"

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<p>%s</p>", body)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

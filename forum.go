package main

import (
	"html/template"
	"net/http"
)

type Post struct {
	Body string
}

var post = &Post{Body: "Nothing to see here"}

func handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, post)
}

func newHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("new.html")
	t.Execute(w, &Post{})
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/new", newHandler)
	http.ListenAndServe(":8080", nil)
}

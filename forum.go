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
	t, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, post)
}

func newHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("new.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, &Post{})
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/new", newHandler)
	http.ListenAndServe(":8080", nil)
}

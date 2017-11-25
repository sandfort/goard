package main

import (
	"html/template"
	"net/http"
)

type Post struct {
	Title string
	Body string
}

var posts []Post

func handler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, posts)
}

func newHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("new.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	body := r.FormValue("body")
	posts = append(posts, Post{Title: title, Body: body})
	http.Redirect(w, r, "/", http.StatusFound)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/new", newHandler)
	http.HandleFunc("/save", saveHandler)
	http.ListenAndServe(":8080", nil)
}

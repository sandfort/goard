package main

import (
	"html/template"
	"net/http"
	"strconv"
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
	http.Redirect(w, r, "/posts", http.StatusFound)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/posts/"):])
	if err != nil {
		http.NotFound(w, r)
		return
	}
	t, err := template.ParseFiles("view.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, posts[id])
}

func main() {
	http.HandleFunc("/posts", handler)
	http.HandleFunc("/posts/new", newHandler)
	http.HandleFunc("/posts/save", saveHandler)
	http.HandleFunc("/posts/", viewHandler)
	http.ListenAndServe(":8080", nil)
}

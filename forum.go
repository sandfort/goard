package main

import (
	"html/template"
	"net/http"
	"strconv"
	"github.com/sandfort/goard/posts"
)

type controller struct {
	store postStore
}

func (c *controller) handler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	posts := c.store.ReadAllPosts()
	t.Execute(w, posts)
}

func (c *controller) newHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("new.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

func (c *controller) saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	body := r.FormValue("body")
	c.store.CreatePost(posts.Post{Title: title, Body: body})
	http.Redirect(w, r, "/posts", http.StatusFound)
}

func (c *controller) viewHandler(w http.ResponseWriter, r *http.Request) {
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
	post, err := c.store.ReadPost(id)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	t.Execute(w, post)
}

type postStore interface {
	ReadPost(id int) (posts.Post, error)
	ReadAllPosts() []posts.Post
	CreatePost(post posts.Post) int
}

func main() {
	ctrl := controller{store: posts.NewMemoryStore()}

	http.HandleFunc("/posts", ctrl.handler)
	http.HandleFunc("/posts/new", ctrl.newHandler)
	http.HandleFunc("/posts/save", ctrl.saveHandler)
	http.HandleFunc("/posts/", ctrl.viewHandler)
	http.ListenAndServe(":8080", nil)
}

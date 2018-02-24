package web

import (
	"github.com/sandfort/goard/core"
	"html/template"
	"net/http"
	"strconv"
)

func NewPostController(store core.PostStore) *controller {
	return &controller{store: store}
}

type controller struct {
	store core.PostStore
}

func (c *controller) Handler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	posts := c.store.ReadAllPosts()
	t.Execute(w, posts)
}

func (c *controller) NewHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("new.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

func (c *controller) SaveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	body := r.FormValue("body")
	c.store.CreatePost(core.Post{Title: title, Body: body})
	http.Redirect(w, r, "/posts", http.StatusFound)
}

func (c *controller) ViewHandler(w http.ResponseWriter, r *http.Request) {
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

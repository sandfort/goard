package web

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/sandfort/goard/core"
)

func NewThreadController(store core.ThreadStore) *controller {
	return &controller{store: store}
}

type controller struct {
	store core.ThreadStore
}

func (c *controller) Handler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	threads := c.store.ReadAllThreads()
	t.Execute(w, threads)
}

func (c *controller) NewHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/new.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

func (c *controller) SaveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	body := r.FormValue("body")
	core.PostNewThread(title, body, c.store)
	http.Redirect(w, r, "/threads", http.StatusFound)
}

func (c *controller) ViewHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/threads/"):])
	if err != nil {
		http.NotFound(w, r)
		return
	}
	t, err := template.ParseFiles("web/view.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	thread, err := c.store.ReadThread(id)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	t.Execute(w, thread)
}

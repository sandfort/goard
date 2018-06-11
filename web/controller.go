package web

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/sandfort/goard/core"
)

func NewThreadController(tstore core.ThreadStore, pstore core.PostStore) *controller {
	return &controller{tstore: tstore, pstore: pstore}
}

type controller struct {
	tstore core.ThreadStore
	pstore core.PostStore
}

func (c *controller) Handler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	threads := c.tstore.ReadAllThreads()
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

	core.PostNewThread(title, body, c.tstore, c.pstore)

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

	thread := core.FetchThreadWithPosts(id, c.tstore, c.pstore)

	if err != nil {
		http.NotFound(w, r)
		return
	}
	t.Execute(w, thread)
}

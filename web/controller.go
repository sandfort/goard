package web

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/sandfort/goard/core"
)

// NewThreadController constructs a controller that handles thread requests.
func NewThreadController(stamper core.Stamper, tstore core.ThreadStore, pstore core.PostStore) *controller {
	return &controller{stamper: stamper, tstore: tstore, pstore: pstore}
}

type controller struct {
	stamper core.Stamper
	tstore  core.ThreadStore
	pstore  core.PostStore
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
	author := r.FormValue("author")

	core.PostNewThread(title, body, author, c.stamper, c.tstore, c.pstore)

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

	thread, err := core.FetchThreadWithPosts(id, c.tstore, c.pstore)

	if err != nil {
		http.NotFound(w, r)
		return
	}
	t.Execute(w, thread)
}

func (c *controller) NewReplyHandler(w http.ResponseWriter, r *http.Request) {
	tid, err := strconv.Atoi(r.URL.Path[len("/threads/reply/"):])

	if err != nil {
		http.NotFound(w, r)
		return
	}

	t, err := template.ParseFiles("web/newreply.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	th, err := core.FetchThreadWithPosts(tid, c.tstore, c.pstore)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	t.Execute(w, th)
}

func (c *controller) SaveReplyHandler(w http.ResponseWriter, r *http.Request) {
	tid, err := strconv.Atoi(r.FormValue("threadId"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	body := r.FormValue("body")
	author := r.FormValue("author")

	core.AddReply(tid, body, author, c.stamper, c.pstore)

	http.Redirect(w, r, fmt.Sprintf("/threads/%d", tid), http.StatusFound)
}

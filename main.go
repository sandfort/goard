package main

import (
	"github.com/sandfort/goard/mem"
	"github.com/sandfort/goard/web"
	"net/http"
)

func main() {
	ctrl := web.NewPostController(mem.NewPostStore())

	http.HandleFunc("/posts", ctrl.Handler)
	http.HandleFunc("/posts/new", ctrl.NewHandler)
	http.HandleFunc("/posts/save", ctrl.SaveHandler)
	http.HandleFunc("/posts/", ctrl.ViewHandler)
	http.ListenAndServe(":8080", nil)
}

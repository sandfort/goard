package main

import (
	"net/http"
	"os"

	"github.com/sandfort/goard/mem"
	"github.com/sandfort/goard/web"
)

func main() {
	ctrl := web.NewPostController(mem.NewPostStore())

	port := os.Getenv("PORT")

	http.HandleFunc("/posts", ctrl.Handler)
	http.HandleFunc("/posts/new", ctrl.NewHandler)
	http.HandleFunc("/posts/save", ctrl.SaveHandler)
	http.HandleFunc("/posts/", ctrl.ViewHandler)
	http.ListenAndServe(":"+port, nil)
}

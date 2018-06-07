package main

import (
	"net/http"
	"os"

	"github.com/sandfort/goard/mem"
	"github.com/sandfort/goard/web"
)

func main() {
	ctrl := web.NewThreadController(mem.NewThreadStore())

	port := os.Getenv("PORT")

	http.HandleFunc("/threads", ctrl.Handler)
	http.HandleFunc("/threads/new", ctrl.NewHandler)
	http.HandleFunc("/threads/save", ctrl.SaveHandler)
	http.HandleFunc("/threads/", ctrl.ViewHandler)
	http.ListenAndServe(":"+port, nil)
}

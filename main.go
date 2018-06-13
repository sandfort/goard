package main

import (
	"net/http"
	"os"

	"github.com/sandfort/goard/core"
	"github.com/sandfort/goard/web"
)

func main() {
	threadStore := core.NewThreadMemoryStore()
	postStore := core.NewPostMemoryStore()

	stamper := core.NewIncrementingStamper()

	ctrl := web.NewThreadController(stamper, threadStore, postStore)

	port := os.Getenv("PORT")

	http.HandleFunc("/threads", ctrl.Handler)
	http.HandleFunc("/threads/new", ctrl.NewHandler)
	http.HandleFunc("/threads/save", ctrl.SaveHandler)
	http.HandleFunc("/threads/", ctrl.ViewHandler)

	http.HandleFunc("/threads/reply/", ctrl.NewReplyHandler)
	http.HandleFunc("/threads/reply/save", ctrl.SaveReplyHandler)

	http.ListenAndServe(":"+port, nil)
}

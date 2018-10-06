package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sandfort/goard/core"
	"github.com/sandfort/goard/db"
	"github.com/sandfort/goard/web"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbFlag := flag.Bool("db", false, "Use database-backed storage")

	flag.Parse()

	var stamper core.Stamper
	var threadStore core.ThreadStore
	var postStore core.PostStore

	if *dbFlag {
		dbUser := os.Getenv("DB_USER")
		dbPassword := os.Getenv("DB_PASSWORD")
		dbName := os.Getenv("DB_NAME")

		dsn := fmt.Sprintf("%s:%s@/%s", dbUser, dbPassword, dbName)
		dbconn, err := sql.Open("mysql", dsn)

		if err != nil {
			log.Fatal("Failed to open a database connection", err)
		}

		threadStore = db.NewThreadDbStore(dbconn)
		postStore = db.NewPostDbStore(dbconn)
	} else {
		threadStore = core.NewThreadMemoryStore()
		postStore = core.NewPostMemoryStore()
	}

	stamper = core.NewTimeStamper()

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

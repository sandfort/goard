package db

import (
	"database/sql"
	"log"

	"github.com/sandfort/goard/core"
)

type PostDbStore struct {
	db *sql.DB
}

func NewPostDbStore(db *sql.DB) core.PostStore {
	return &PostDbStore{db: db}
}

func (store *PostDbStore) CreatePost(post core.Post) int {
	result, err := store.db.Exec("insert into post (thread_id, body, author, stamp) values (?, ?, ?, ?)",
		post.ThreadId, post.Body, post.Author, post.Stamp)
	if err != nil {
		log.Fatal(err)
	}
	id, err := result.LastInsertId()
	return int(id)
}

func (store *PostDbStore) ReadPost(id int) (core.Post, error) {
	row := store.db.QueryRow("select id, thread_id, body, author, stamp from post where id = ?", id)
	post := core.Post{}
	err := row.Scan(&post.Id, &post.ThreadId, &post.Body, &post.Author, &post.Stamp)

	if err != nil {
		return core.Post{}, err
	}

	return post, nil
}

func (store *PostDbStore) ReadByThreadId(tid int) []core.Post {
	rows, _ := store.db.Query("select id, thread_id, body, author, stamp from post where thread_id = ?", tid)

	posts := make([]core.Post, 0)
	defer rows.Close()
	for rows.Next() {
		post := core.Post{}
		rows.Scan(&post.Id, &post.ThreadId, &post.Body, &post.Author, &post.Stamp)
		posts = append(posts, post)
	}
	return posts
}

func (store *PostDbStore) DeletePost(id int) {
	store.db.Exec("delete from post where id = ?", id)
}

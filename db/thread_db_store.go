package db

import (
	"database/sql"

	"github.com/sandfort/goard/core"
)

type ThreadDbStore struct {
	db *sql.DB
}

func NewThreadDbStore(db *sql.DB) core.ThreadStore {
	return &ThreadDbStore{db: db}
}

func (store *ThreadDbStore) ReadThread(id int) (core.Thread, error) {
	row := store.db.QueryRow("select id, title from thread where id = ?", id)

	thread := core.Thread{}
	err := row.Scan(&thread.Id, &thread.Title)

	if err != nil {
		return core.Thread{}, err
	}

	return thread, nil
}

func (store *ThreadDbStore) ReadAllThreads() []core.Thread {
	rows, _ := store.db.Query("select id, title from thread")

	threads := make([]core.Thread, 0)
	defer rows.Close()
	for rows.Next() {
		thread := core.Thread{}
		rows.Scan(&thread.Id, &thread.Title)
		threads = append(threads, thread)
	}
	return threads
}

func (store *ThreadDbStore) CreateThread(thread core.Thread) int {
	result, _ := store.db.Exec("insert into thread (title) values (?)", thread.Title)
	id, _ := result.LastInsertId()
	return int(id)
}

func (store *ThreadDbStore) DeleteThread(id int) {
	store.db.Exec("delete from thread where id = ?", id)
}

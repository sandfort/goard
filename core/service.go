package core

import (
	"sort"
)

// PostNewThread posts a new thread with the given title and body to store, returning its ID.
func PostNewThread(title string, body string, author string, stamper Stamper, tstore ThreadStore, pstore PostStore) int {
	tid := tstore.CreateThread(Thread{Title: title})
	pstore.CreatePost(Post{Body: body, Author: author, ThreadID: tid, Stamp: stamper.Stamp()})

	return tid
}

// FetchThreadWithPosts returns a thread and all posts associated with it.
func FetchThreadWithPosts(id int, tstore ThreadStore, pstore PostStore) (ThreadWithPosts, error) {
	t, err := tstore.ReadThread(id)

	if err != nil {
		return ThreadWithPosts{}, err
	}

	ps := pstore.ReadByThreadId(id)

	sort.Slice(ps, func(i, j int) bool {
		return ps[i].Stamp < ps[j].Stamp
	})

	return ThreadWithPosts{ThreadId: t.Id, Title: t.Title, Posts: ps}, nil
}

// AddReply saves a new reply to the thread with the given ID.
func AddReply(tid int, body string, author string, stamper Stamper, pstore PostStore) {
	pstore.CreatePost(Post{ThreadID: tid, Body: body, Author: author, Stamp: stamper.Stamp()})
}

package core

// PostNewThread posts a new thread with the given title and body to store, returning its ID.
func PostNewThread(title string, body string, tstore ThreadStore, pstore PostStore) int {
	tid := tstore.CreateThread(Thread{Title: title})
	pstore.CreatePost(Post{Body: body, ThreadId: tid})

	return tid
}

// FetchThreadWithPosts returns a thread and all posts associated with it.
func FetchThreadWithPosts(id int, tstore ThreadStore, pstore PostStore) ThreadWithPosts {
	t, _ := tstore.ReadThread(id)
	ps := pstore.ReadByThreadId(id)
	return ThreadWithPosts{ThreadId: t.Id, Title: t.Title, Posts: ps}
}

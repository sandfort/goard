package core

// PostNewThread a new thread with the given title and body to store, returning its ID.
func PostNewThread(title string, body string, tstore ThreadStore, pstore PostStore) int {
	tid := tstore.CreateThread(Thread{Title: title, Body: body})
	pstore.CreatePost(Post{Body: body, ThreadId: tid})

	return tid
}

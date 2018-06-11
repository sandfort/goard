package core

// PostNewThread posts a new thread with the given title and body to store, returning its ID.
func PostNewThread(title string, body string, tstore ThreadStore, pstore PostStore) int {
	tid := tstore.CreateThread(Thread{Title: title})
	pstore.CreatePost(Post{Body: body, ThreadId: tid})

	return tid
}

// FetchThreadWithPosts returns a thread and all posts associated with it.
func FetchThreadWithPosts(id int, tstore ThreadStore, pstore PostStore) (ThreadWithPosts, error) {
	t, err := tstore.ReadThread(id)

	if err != nil {
		return ThreadWithPosts{}, err
	}

	ps := pstore.ReadByThreadId(id)
	return ThreadWithPosts{ThreadId: t.Id, Title: t.Title, Posts: ps}, nil
}

func AddReply(tid int, body string, pstore PostStore) {
	pstore.CreatePost(Post{ThreadId: tid, Body: body})
}

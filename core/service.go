package core

// PostNewThread a new thread with the given title and body to store, returning its ID.
func PostNewThread(title string, body string, store ThreadStore) int {
	return store.CreateThread(Thread{Title: title, Body: body})
}

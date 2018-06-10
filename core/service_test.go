package core

import "testing"

func TestPostNewThread(t *testing.T) {
	threadStore := NewThreadMemoryStore()
	postStore := NewPostMemoryStore()

	title := "the title"
	body := "the body"

	PostNewThread(title, body, threadStore, postStore)

	threads := threadStore.ReadAllThreads()

	if len(threads) != 1 {
		t.Errorf("Expected 1 thread but got %d", len(threads))
	}

	createdThread := threads[0]

	if createdThread.Title != title {
		t.Errorf("Title should have been %q but was %q", title, createdThread.Title)
	}

	posts := postStore.ReadByThreadId(createdThread.Id)

	if len(posts) != 1 {
		t.Errorf("Expected thread to have 1 post but had %d", len(posts))
	}

	createdPost := posts[0]

	if createdPost.Body != body {
		t.Errorf("Body should have been %q but was %q", body, createdPost.Body)
	}
}

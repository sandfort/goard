package core

import "testing"

func TestPostNewThread(t *testing.T) {
	threadStore := NewThreadMemoryStore()

	title := "the title"
	body := "the body"

	PostNewThread(title, body, threadStore)

	threads := threadStore.ReadAllThreads()

	if len(threads) != 1 {
		t.Errorf("Expected 1 but got %d", len(threads))
	}

	createdThread := threads[0]

	if createdThread.Title != title {
		t.Errorf("Title should have been %q but was %q", title, createdThread.Title)
	}

	if createdThread.Body != body {
		t.Errorf("Body should have been %q but was %q", body, createdThread.Body)
	}
}

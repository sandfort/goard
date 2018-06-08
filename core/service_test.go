package core

import "testing"

func TestPostNewThread(t *testing.T) {
	threadStore := NewThreadMemoryStore()

	PostNewThread("the title", "the body", threadStore)

	threads := threadStore.ReadAllThreads()

	if len(threads) != 1 {
		t.Errorf("Expected 1 but got %d", len(threads))
	}

	createdThread := threads[0]

	if createdThread.Title != "the title" {
		t.Error("Title did not match")
	}

	if createdThread.Body != "the body" {
		t.Error("Body did not match")
	}
}

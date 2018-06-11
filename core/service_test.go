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

func TestFetchThreadWithPosts(t *testing.T) {
	threadStore := NewThreadMemoryStore()
	postStore := NewPostMemoryStore()

	title := "the title"
	body := "the body"

	tid := PostNewThread(title, body, threadStore, postStore)

	th, _ := FetchThreadWithPosts(tid, threadStore, postStore)

	if th.ThreadId != tid {
		t.Errorf("Expected to find thread with id %d but got %d", tid, th.ThreadId)
	}

	if th.Title != title {
		t.Errorf("Expected thread to have title %q but got %q", title, th.Title)
	}

	if len(th.Posts) != 1 {
		t.Errorf("Expected thread to have 1 post but found %d", len(th.Posts))
	}
}

func TestFetchThreadWithPostsReturnsErrorWhenThreadDoesNotExist(t *testing.T) {
	threadStore := NewThreadMemoryStore()
	postStore := NewPostMemoryStore()

	_, err := FetchThreadWithPosts(1, threadStore, postStore)

	if err == nil {
		t.Error("Expected method to return error")
	}
}

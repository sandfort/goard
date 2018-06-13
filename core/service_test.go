package core

import "testing"

func TestPostNewThread(t *testing.T) {
	threadStore := NewThreadMemoryStore()
	postStore := NewPostMemoryStore()
	stamper := NewIncrementingStamper()

	title := "the title"
	body := "the body"
	author := "the author"

	PostNewThread(title, body, author, stamper, threadStore, postStore)

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

	if createdPost.Author != author {
		t.Errorf("Author should have been %q but was %q", author, createdPost.Author)
	}
}

func TestFetchThreadWithPosts(t *testing.T) {
	threadStore := NewThreadMemoryStore()
	postStore := NewPostMemoryStore()
	stamper := NewIncrementingStamper()

	title := "the title"
	body := "the body"
	author := "the author"

	tid := PostNewThread(title, body, author, stamper, threadStore, postStore)

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

func TestFetchThreadWithPostsReturnsPostsInOrder(t *testing.T) {
	threadStore := NewThreadMemoryStore()
	postStore := NewPostMemoryStore()
	stamper := NewIncrementingStamper()

	title := "thread title"
	author := "author"
	body1 := "first post"
	body2 := "second post"
	body3 := "third post"
	body4 := "fourth post"

	tid := PostNewThread(title, body1, author, stamper, threadStore, postStore)
	AddReply(tid, body2, author, stamper, postStore)
	AddReply(tid, body3, author, stamper, postStore)
	AddReply(tid, body4, author, stamper, postStore)

	th, _ := FetchThreadWithPosts(tid, threadStore, postStore)

	if th.Posts[0].Body != body1 {
		t.Errorf("Expected first post body to be %q but was %q", body1, th.Posts[0].Body)
	}

	if th.Posts[1].Body != body2 {
		t.Errorf("Expected second post body to be %q but was %q", body2, th.Posts[1].Body)
	}

	if th.Posts[2].Body != body3 {
		t.Errorf("Expected third post body to be %q but was %q", body3, th.Posts[2].Body)
	}

	if th.Posts[3].Body != body4 {
		t.Errorf("Expected fourth post body to be %q but was %q", body4, th.Posts[3].Body)
	}
}

func TestAddReply(t *testing.T) {
	title := "thread title"
	post1 := "first post"
	post2 := "second post"
	author := "the author"

	tstore := NewThreadMemoryStore()
	pstore := NewPostMemoryStore()
	stamper := NewIncrementingStamper()

	tid := PostNewThread(title, post1, author, stamper, tstore, pstore)

	AddReply(tid, post2, author, stamper, pstore)

	th, _ := FetchThreadWithPosts(tid, tstore, pstore)

	if len(th.Posts) != 2 {
		t.Errorf("Expected thread to have 2 posts but found %d", len(th.Posts))
	}

	for _, p := range th.Posts {
		if p.Author != author {
			t.Errorf("Expected post to have author %q but was %q", author, p.Author)
		}
	}
}

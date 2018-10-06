package core

import "testing"

func NewPostStoreContract(pstore PostStore, tstore ThreadStore) *fixture {
	return &fixture{pstore: pstore, tstore: tstore}
}

type fixture struct {
	pstore PostStore
	tstore ThreadStore
}

type PostStoreContract interface {
	Verify(t *testing.T)
	createAndReadPostContract(t *testing.T)
	createAndReadMultiplePostsContract(t *testing.T)
	readPostReturnsErrorWhenIdDoesNotExistContract(t *testing.T)
	readByThreadIdContract(t *testing.T)
}

func (f *fixture) Verify(t *testing.T) {
	f.createAndReadPostContract(t)
	f.createAndReadMultiplePostsContract(t)
	f.readPostReturnsErrorWhenIdDoesNotExistContract(t)
	f.readByThreadIdContract(t)
}

func (f *fixture) createAndReadPostContract(t *testing.T) {
	tid := f.tstore.CreateThread(Thread{})
	defer f.tstore.DeleteThread(tid)

	body := "The Body"

	id := f.pstore.CreatePost(Post{Body: body, ThreadId: tid})
	defer f.pstore.DeletePost(id)

	result, _ := f.pstore.ReadPost(id)
	if result.Body != body {
		t.Errorf("Was expecting body to be %q but was %q", body, result.Body)
	}
}

func (f *fixture) createAndReadMultiplePostsContract(t *testing.T) {
	tid := f.tstore.CreateThread(Thread{})
	defer f.tstore.DeleteThread(tid)

	body1 := "first"
	body2 := "second"

	id1 := f.pstore.CreatePost(Post{Body: body1, ThreadId: tid, Stamp: 0})
	defer f.pstore.DeletePost(id1)
	id2 := f.pstore.CreatePost(Post{Body: body2, ThreadId: tid, Stamp: 1})
	defer f.pstore.DeletePost(id2)

	post1, _ := f.pstore.ReadPost(id1)
	post2, _ := f.pstore.ReadPost(id2)

	if post1.Body != body1 {
		t.Errorf("Expected body to be %q but was %q", body1, post1.Body)
	}

	if post2.Body != body2 {
		t.Errorf("Expected body to be %q but was %q", body2, post2.Body)
	}
}

func (f *fixture) readPostReturnsErrorWhenIdDoesNotExistContract(t *testing.T) {
	_, err := f.pstore.ReadPost(1)

	if err == nil {
		t.Error("Expected error")
	}
}

func (f *fixture) readByThreadIdContract(t *testing.T) {
	tid1 := f.tstore.CreateThread(Thread{})
	defer f.tstore.DeleteThread(tid1)
	tid2 := f.tstore.CreateThread(Thread{})
	defer f.tstore.DeleteThread(tid2)

	id1 := f.pstore.CreatePost(Post{ThreadId: tid1, Stamp: 0})
	defer f.pstore.DeletePost(id1)
	id2 := f.pstore.CreatePost(Post{ThreadId: tid2, Stamp: 1})
	defer f.pstore.DeletePost(id2)
	id3 := f.pstore.CreatePost(Post{ThreadId: tid2, Stamp: 2})
	defer f.pstore.DeletePost(id3)

	t1Posts := f.pstore.ReadByThreadId(tid1)

	if len(t1Posts) != 1 {
		t.Errorf("Expected thread to have 1 post but found %d", len(t1Posts))
	}

	if t1Posts[0].Id != id1 {
		t.Errorf("Expected to find post with ID %d but found ID %d", id1, t1Posts[0].Id)
	}
}

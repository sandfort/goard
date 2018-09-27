package core

// Post is the basic type representing any post in a thread by a user.
type Post struct {
	ID       int
	ThreadID int
	Body     string
	Author   string
	Stamp    int
}

// PostStore is the generic interface for storing and retrieving Post objects.
type PostStore interface {
	CreatePost(post Post) int
	ReadPost(id int) (Post, error)
	ReadByThreadId(tid int) []Post
}

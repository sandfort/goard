package core

type Post struct {
	Id       int
	ThreadId int
	Body     string
	Author   string
	Stamp    int
}

type PostStore interface {
	CreatePost(post Post) int
	ReadPost(id int) (Post, error)
	ReadByThreadId(tid int) []Post
}

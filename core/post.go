package core

type Post struct {
	Id    int
	Title string
	Body  string
}

type PostStore interface {
	ReadPost(id int) (Post, error)
	ReadAllPosts() []Post
	CreatePost(post Post) int
}

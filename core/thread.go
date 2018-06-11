package core

type Thread struct {
	Id    int
	Title string
}

type ThreadWithPosts struct {
	ThreadId int
	Title    string
	Posts    []Post
}

type ThreadStore interface {
	ReadThread(id int) (Thread, error)
	ReadAllThreads() []Thread
	CreateThread(thread Thread) int
}

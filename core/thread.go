package core

type Thread struct {
	Id    int
	Title string
	Body  string
}

type ThreadStore interface {
	ReadThread(id int) (Thread, error)
	ReadAllThreads() []Thread
	CreateThread(thread Thread) int
}

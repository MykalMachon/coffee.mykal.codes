package models

type Post struct {
	ID       int
	Title    string
	Body     string
	ImageURL string
}

type PostService struct{}

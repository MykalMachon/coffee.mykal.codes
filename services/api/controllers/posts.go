package controllers

import (
	"fmt"
	"net/http"

	"github.com/mykalmachon/coffee.mykal.codes/api/models"
)

type PostController struct {
	PostService *models.PostService
}

func (pc *PostController) GetPost(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	// TODO: get post and it's data
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello: you requested post with id = %s", id)
}

func (pc *PostController) GetPosts(w http.ResponseWriter, r *http.Request) {
	// TODO: get first page of posts
	// TODO: get optional page query param to get specific page
	// TODO: allow filtering by tags from query params
	w.Write([]byte("Hello: you requested all posts"))
}

func (pc *PostController) CreatePost(w http.ResponseWriter, r *http.Request) {
	// TODO: get post data from request and save it
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Hello: you created a new post"))
}

func (pc *PostController) UpdatePost(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	// TODO: update post and it's data
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello: you updated post with id = " + id))
}

func (pc *PostController) DeletePost(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "you requested to delete %s", id)
}

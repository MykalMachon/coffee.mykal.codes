package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/rs/cors" // cors goodies
)

func main() {

	portStr := os.Getenv("PORT")
	port := 8080 // default port
	if portStr != "" {
		// Convert the port string to an integer
		port, _ = strconv.Atoi(portStr)
	}
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	router.HandleFunc("GET /healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Service Healthy")
	})

	router.HandleFunc("GET /posts/", func(w http.ResponseWriter, r *http.Request) {
		// TODO: get first page of posts
		// TODO: get optional page query param to get specific page
		// TODO: allow filtering by tags from query params
		w.Write([]byte("Hello: you requested all posts"))
	})

	router.HandleFunc("POST /posts/", func(w http.ResponseWriter, r *http.Request) {
		// TODO: get post data from request and save it
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Hello: you created a new post"))
	})

	router.HandleFunc("GET /posts/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		// TODO: get post and it's data
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello: you requested post with id = " + id))
	})

	router.HandleFunc("PUT /posts/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		// TODO: update post and it's data
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello: you updated post with id = " + id))
	})

	c := cors.New(cors.Options{})

	server := http.Server{
		Addr:    fmt.Sprintf("[::]:%d", port),
		Handler: c.Handler(router),
	}

	log.Printf("Starting server on port [::]:%d", port)
	server.ListenAndServe()
}

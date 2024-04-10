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
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Hello World"))
		} else {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "Page not found")
		}
	})

	// * HEALTH CHECK ROUTES
	router.HandleFunc("GET /healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprint(w, "Service Healthy")
	})

	// * AUTH ROUTES
	router.HandleFunc("GET /auth/", func(w http.ResponseWriter, r *http.Request) {
		// TODO: return authentication status
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Not authorized")
	})

	router.HandleFunc("POST /auth/signup", func(w http.ResponseWriter, r *http.Request) {
		// TODO: get name/username/password
		// TODO: check if accepting signups flag is active (if one user is there)
		// TODO: hash password
		// TODO: create user object
		// TODO: create a session, create a cookie and attach it to w
	})

	router.HandleFunc("POST /auth/login", func(w http.ResponseWriter, r *http.Request) {
		// TODO: get username/password.
		// TODO: lookup username; return 401 if it no user exists
		// TODO: hash password and check if password hashes match; return 401 if not
		// TODO: if valid login, create a cookie and attach it to w
		// TODO: if valid login, and mode=API return api key (maybe?)
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Not authorized")
	})

	router.HandleFunc("GET /auth/logout", func(w http.ResponseWriter, r *http.Request) {
		// TODO: clear the cookie from the request
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Logged out")
	})

	// * POSTS ROUTES
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

	portStr := os.Getenv("PORT")
	port := 8080 // default port
	if portStr != "" {
		// Convert the port string to an integer
		port, _ = strconv.Atoi(portStr)
	}

	server := http.Server{
		Addr:    fmt.Sprintf("[::]:%d", port),
		Handler: c.Handler(router),
	}

	log.Printf("Starting server on port [::]:%d", port)
	server.ListenAndServe()
}

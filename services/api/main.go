package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/mykalmachon/coffee.mykal.codes/api/controllers"
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

	// TODO: fix bad 404 routes with middleware
	// right now a call to /posts/123/hello/world is not a 404. it defaults to /posts/123
	// need to fix this; should be able to do it with DI / middleware approach 🤔

	// * META ROUTES
	metaController := controllers.MetaController{}
	router.HandleFunc("GET /meta/healthcheck", metaController.Healtcheck)

	// * AUTH ROUTES
	authController := controllers.AuthController{}
	router.HandleFunc("GET /auth/", authController.Status)
	router.HandleFunc("POST /auth/signup", authController.Signup)
	router.HandleFunc("POST /auth/login", authController.Login)
	router.HandleFunc("GET /auth/logout", authController.Logout)

	// * POSTS ROUTES
	postController := controllers.PostController{}
	router.HandleFunc("GET /posts/", postController.GetPosts)
	router.HandleFunc("POST /posts/", postController.CreatePost)
	router.HandleFunc("GET /posts/{id}", postController.GetPost)
	router.HandleFunc("PUT /posts/{id}", postController.UpdatePost)
	router.HandleFunc("DELETE /posts/{id}", postController.DeletePost)

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

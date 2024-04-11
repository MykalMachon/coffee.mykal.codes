package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5" // default handler still sucks at 404/middleware
	"github.com/rs/cors"       // used for managing cors and other goodies

	"github.com/mykalmachon/coffee.mykal.codes/api/controllers"
)

func main() {
	router := chi.NewRouter()

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

	// * ERRORS
	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Route not found", http.StatusNotFound)
	})

	portStr := os.Getenv("PORT")
	port := 8080 // default port
	if portStr != "" {
		port, _ = strconv.Atoi(portStr)
	}

	log.Printf("Starting server on port [::]:%d", port)

	corsHandler := cors.Default().Handler(router) // wraps chi Mux in Cors Handler
	http.ListenAndServe(fmt.Sprintf("[::]:%d", port), corsHandler)
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5" // default handler still sucks at 404/middleware
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv" // for local dev
	"github.com/rs/cors"       // used for managing cors and other goodies
	"gorm.io/driver/postgres"
	"gorm.io/gorm" // ORM for database

	"github.com/mykalmachon/coffee.mykal.codes/api/controllers"
	"github.com/mykalmachon/coffee.mykal.codes/api/models"
)

func OpenDatabaseConnection() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: os.Getenv("DATABASE_URL"),
	}), &gorm.Config{
		TranslateError: true,
	})
	return db, err
}

func RunDatabaseMigrations(db *gorm.DB, models ...interface{}) error {
	err := db.AutoMigrate(models)
	return err
}

func main() {
	godotenv.Load()

	// * DATABASE SETUP
	db, dbErr := OpenDatabaseConnection()

	if dbErr != nil {
		log.Printf("Cannot connect to database, shutting down: %s", dbErr.Error())
		os.Exit(1)
	}

	dbErr = db.AutoMigrate(models.User{}, models.Post{})

	if dbErr != nil {
		log.Printf("Could not run datbase migrations, shutting down %s", dbErr.Error())
		os.Exit(1)
	}

	// * MODEL SETUP
	userService := models.UserService{DB: db}
	postServices := models.PostService{DB: db}

	// * MIDDLEWARE MANAGERS
	authMiddleware := controllers.AuthMiddleware{UserService: &userService}

	// * ROUTER SETUP AND MIDDLEWARE
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(authMiddleware.SetAuthUser)

	// * META ROUTES
	metaController := controllers.MetaController{}
	router.HandleFunc("GET /meta/healthcheck", metaController.Healtcheck)

	// * AUTH ROUTES
	// see OAuth examples here: https://github.com/go-chi/oauth/blob/master/example/authserver/main.go
	authController := controllers.AuthController{UserSerivce: &userService}
	router.HandleFunc("GET /auth/", authController.Status)
	router.HandleFunc("POST /auth/signup", authController.Signup)
	router.HandleFunc("POST /auth/login", authController.Login)

	// * POSTS ROUTES
	postController := controllers.PostController{PostService: &postServices}
	router.HandleFunc("GET /posts/", postController.GetPosts)
	router.HandleFunc("POST /posts/", postController.CreatePost)
	router.HandleFunc("GET /posts/{id}", postController.GetPost)
	router.HandleFunc("PUT /posts/{id}", postController.UpdatePost)
	router.HandleFunc("DELETE /posts/{id}", postController.DeletePost)

	// * ERRORS
	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Route not found", http.StatusNotFound)
	})

	// * SERVER SETUP
	portStr := os.Getenv("PORT")
	port := 8080 // default port
	if portStr != "" {
		port, _ = strconv.Atoi(portStr)
	}

	log.Printf("Starting server on port [::]:%d", port)
	corsHandler := cors.Default().Handler(router) // TODO: move to official chi cors middleware
	http.ListenAndServe(fmt.Sprintf("[::]:%d", port), corsHandler)
}

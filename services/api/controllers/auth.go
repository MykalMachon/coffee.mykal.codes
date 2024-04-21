package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mykalmachon/coffee.mykal.codes/api/models"
	"gorm.io/gorm"
)

type AuthController struct {
	UserSerivce *models.UserService
}

// * UTILITY FUNCTIONS
func createAuthToken(user models.User) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return "", errors.New("no jwt secret available")
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"name":  user.Name,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := t.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", fmt.Errorf("create auth token: %w", err)
	}
	return tokenString, nil
}

func validateAuthToken(token string) (bool, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return false, errors.New("no jwt secret available")
	}

	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Verify the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret key used for signing
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return false, fmt.Errorf("parse auth token: %w", err)
	}

	// Check if the token is valid
	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		exp := claims["exp"].(float64)
		expTime := time.Unix(int64(exp), 0)
		if time.Now().After(expTime) {
			return false, nil
		}
		return true, nil
	}

	return false, nil
}

// * MIDDLEWARES

type key int

const userKey key = iota

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get user from auth
		// var user models.User
		// authToken := r.Header.Get("Authorization")

		// if authToken != "" {
		// 	println("auth token found in request")
		// } else {
		// 	println("auth token not found in request")
		// }

		ctx := context.WithValue(r.Context(), userKey, nil)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// * CONTROLLER ROUTES
func (ac *AuthController) Status(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	valid, err := validateAuthToken(token)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}
	if !valid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}
	fmt.Fprint(w, "Token is valid!")
}

func (ac *AuthController) Signup(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")
	passwordConf := r.FormValue("passwordConfirmation")

	if password != passwordConf {
		http.Error(w, "Passwords do not match.", http.StatusBadRequest)
		return
	}

	user, err := ac.UserSerivce.CreateUser(name, email, password)

	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			http.Error(w, "User already exists with this email", http.StatusBadRequest)
		} else {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User created: %+v", user)
}

func (ac *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	type ResponseData struct {
		Message string `json:"message"`
		Token   string `json:"token"`
	}

	email := r.FormValue("email")
	password := r.FormValue("password")
	user, err := ac.UserSerivce.Authenticate(email, password)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(ResponseData{
			Message: "Invalid login. Please try again!",
			Token:   "",
		})
		return
	}

	token, err := createAuthToken(*user)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ResponseData{
			Message: "Something went wrong. Please try again!",
			Token:   "",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ResponseData{
		Message: "Login successful!",
		Token:   token,
	})
}

package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/mykalmachon/coffee.mykal.codes/api/models"
	"gorm.io/gorm"
)

type AuthController struct {
	UserSerivce *models.UserService
}

func (ac *AuthController) Status(w http.ResponseWriter, r *http.Request) {
	// TODO: return authentication status
	w.WriteHeader(http.StatusUnauthorized)
	fmt.Fprint(w, "Not authorized")
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
	email := r.FormValue("email")
	password := r.FormValue("password")

	_, err := ac.UserSerivce.Authenticate(email, password)

	if err != nil {
		http.Error(w, "Invalid login. Try again.", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User login successful")
}

func (ac *AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	// TODO: clear the cookie from the request
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Logged out")
}

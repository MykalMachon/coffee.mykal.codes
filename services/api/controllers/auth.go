package controllers

import (
	"fmt"
	"net/http"
)

type AuthController struct{}

func (ac *AuthController) Status(w http.ResponseWriter, r *http.Request) {
	// TODO: return authentication status
	w.WriteHeader(http.StatusUnauthorized)
	fmt.Fprint(w, "Not authorized")
}

func (ac *AuthController) Signup(w http.ResponseWriter, r *http.Request) {
	// TODO: get name/username/password
	// TODO: check if accepting signups flag is active (if one user is there)
	// TODO: hash password
	// TODO: create user object
	// TODO: create a session, create a cookie and attach it to w
	w.WriteHeader(http.StatusUnauthorized)
	fmt.Fprint(w, "Signup is not allowed at this time")
}

func (ac *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	// TODO: get username/password.
	// TODO: lookup username; return 401 if it no user exists
	// TODO: hash password and check if password hashes match; return 401 if not
	// TODO: if valid login, create a cookie and attach it to w
	// TODO: if valid login, and mode=API return api key (maybe?)
	w.WriteHeader(http.StatusUnauthorized)
	fmt.Fprint(w, "Not authorized")
}

func (ac *AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	// TODO: clear the cookie from the request
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Logged out")
}

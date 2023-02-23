package login

import (
	"net/http"
	"time"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {}
func (auth *Authentication) Login(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	if username == "" || password == "" {
		http.Error(w, "missing requorwd fields", http.StatusBadRequest)
		return
	}
	token, err := auth.AuthenticateUser(username, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	accessTokenCookie := http.Cookie{
		Name:     "access_token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24 * 30),
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, &accessTokenCookie)
	w.WriteHeader(http.StatusOK)
}
func (Auth *Authentication) Logout() {}

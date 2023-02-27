package login

import (
	"net/http"
	"time"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	action := r.URL.Query().Get("action")

	if action == "logout" {
		logout(w, r, "access_token")
		return
	}
	if action == "login" {
		login(w, r)
		return
	}
}
func (auth *Authentication) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
func login(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	if username == "" || password == "" {
		http.Error(w, "missing request fields", http.StatusBadRequest)
		return
	}
	token, err := AuthenticateUser(username, password)
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
func logout(w http.ResponseWriter, r *http.Request, cookieName string) {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	cookie.Expires = time.Now().AddDate(0, 0, -1)
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
}

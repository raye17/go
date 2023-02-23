package router

import (
	"git.inspur.com/szsciit/cnos/adapter/pkg/login"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Router() {
	r := mux.NewRouter()
	r.HandleFunc("/apis/cnos.inspur.com/v1/user/login", login.LoginHandler).Methods("POST")
	if err := http.ListenAndServe(":8099", r); err != nil {
		log.Fatalln(err)
	}
}

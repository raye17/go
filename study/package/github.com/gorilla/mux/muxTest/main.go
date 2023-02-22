package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"study/package/github.com/gorilla/mux/muxTest/model"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", model.BooksHandler)
	r.HandleFunc("/books/{id}", model.BookHandle)
	http.Handle("/", r)
	log.Fatalln(http.ListenAndServe(":8077", nil))
}

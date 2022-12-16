package model

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type Book struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Authors     []string `json:"authors"`
	Press       string   `json:"press"`
	PublishedAt string   `json:"published_at"`
}

var (
	mapBooks map[string]*Book
	slcBooks []*Book
)

func init() {
	mapBooks = make(map[string]*Book)
	slcBooks = make([]*Book, 0, 1)
	data, err := ioutil.ReadFile("data/book.json")
	if err != nil {
		log.Fatalln("failed to read book:", err)
	}
	err = json.Unmarshal(data, &slcBooks)
	if err != nil {
		log.Fatalln("failed to unmarshal data:", err)
	}
	for _, book := range slcBooks {
		mapBooks[book.Id] = book
	}
}
func BooksHandler(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	enc.Encode(slcBooks)
}
func BookHandle(w http.ResponseWriter, r *http.Request) {
	book, ok := mapBooks[mux.Vars(r)["id"]]
	if !ok {
		http.NotFound(w, r)
		return
	}
	enc := json.NewEncoder(w)
	enc.Encode(book)
}

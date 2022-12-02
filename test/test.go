package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println(time.Now())
	fmt.Println(time.Now().Format("2006-01-02 15:03:03"))
	http.HandleFunc("/login", login)
	http.ListenAndServe("localhost:9999", nil)

}
func login(w http.ResponseWriter, r *http.Request) {
	r.Body.Close()
}

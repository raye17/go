package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type TestHandler struct {
	str string
}

var i, j = 0, 0

func SayHello(w http.ResponseWriter, r *http.Request) {
	i++
	log.Println("SayHello func is running...", i)
	w.Write([]byte(("hello,this is SayHello func...,this is" + strconv.Itoa(i) + "time")))
}
func (t *TestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	j++
	log.Println("testHandler serverHttp...")
	w.Write([]byte("testHandler.ServerHttp " + strconv.Itoa(j) + "time"))
}
func main() {
	fmt.Println("main start...")
	http.Handle("/", &TestHandler{"hello,my name is sxy"})
	http.HandleFunc("/test", SayHello)
	fmt.Println("listen...")
	http.ListenAndServe("127.0.0.1:8899", nil)
}

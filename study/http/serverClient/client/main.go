package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	client := &http.Client{}
	rootUrl := "http://127.0.0.1:8899/"
	HelloClient(client, rootUrl+"test", "GET")
}
func HelloClient(client *http.Client, url, method string) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	rep, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("client do...")
	data, err := ioutil.ReadAll(rep.Body)
	rep.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s\n", string(data))
}

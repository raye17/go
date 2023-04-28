package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	setupLogger()
	simpleHttpGet("www.baidu.com")
	simpleHttpGet("https://www.baidu.com")
}
func setupLogger() {
	logFile, err := os.OpenFile("./test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(logFile)
}
func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("error fetching url %s:%s", url, err.Error())
	} else {
		log.Printf("status Code for %s:%s", url, resp.Status)
		resp.Body.Close()
	}
	return
}

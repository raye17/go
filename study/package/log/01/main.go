package main

import "log"

func main() {
	log.Println("001")
	log.Println("002")
}
func init() {
	log.SetPrefix("[user]")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

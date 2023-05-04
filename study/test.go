package main

import "log"

func main() {
	log.SetPrefix("pre-login:")
	log.Println("login error!")
	log.SetPrefix("pre-logout:")
	log.Println("sss")
	log.Printf("%s", log.Prefix())
}

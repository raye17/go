package main

import (
	"fmt"
	"log"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("this is panic recover")
			log.Println("recover", err)
		}
	}()
	panic("test11")
	log.Println("ssssss")
}

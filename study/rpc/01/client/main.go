package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Params struct {
	Width, Height int
}

func main() {
	conn, err := rpc.DialHTTP("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	ret := 0
	err2 := conn.Call("Rect.Area", Params{12, 3}, &ret)
	if err2 != nil {
		log.Panic(err)
	}
	fmt.Println("area:", ret)
}

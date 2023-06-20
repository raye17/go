package main

import (
	"fmt"
	"log"
	"net/http"
	"net/rpc"
)

type Params struct {
	Width, Height int
}
type Rect struct {
}

func (r *Rect) Area(p Params, ret *int) error {
	*ret = p.Width * p.Height
	return nil
}
func main() {
	rect := &Rect{}
	rpc.Register(rect)
	rpc.HandleHTTP()
	fmt.Println("ok")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panic(err)
	}
}

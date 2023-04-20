package main

import "fmt"

type secret struct {
	name string
}
type options struct {
	credential *secret
	cab        []byte
	inSecure   bool
	headers    map[string]string
}

func main() {
	opts := &options{}
	fmt.Println("h", opts.headers)
	fmt.Println("c", opts.cab)
	fmt.Println("i", opts.inSecure)
	fmt.Println("cre", opts.credential.name)
}

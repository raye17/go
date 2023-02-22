package main

import (
	"context"
	"fmt"
)

type favContextKey string

func main() {
	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}
	f1 := func(ctx context.Context, k string) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}
	k := favContextKey("color")
	ctx01 := context.WithValue(context.Background(), "language", "Go")
	ctx02 := context.WithValue(context.Background(), k, "Red")
	f1(ctx01, "language")
	f(ctx02, favContextKey("color"))
}

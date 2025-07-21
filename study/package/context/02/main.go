package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	Test04()
}
func Test04() {
	deadline := time.Now().Add(3 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()
	select {
	case <-ctx.Done():
		fmt.Println("ctx.done....")
	case <-time.After(5 * time.Second):
		fmt.Println("5 s...")

	}
}

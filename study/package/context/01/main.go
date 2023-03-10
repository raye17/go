package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

const keyID = "id"

func init() {
	rand.Seed(time.Now().UnixNano())
	context.Canceled = errors.New("sss")

}
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	for i := 0; i < 5; i++ {
		go func(i int) {
			fmt.Println(i)
			select {
			case <-ctx.Done():
				fmt.Printf("%d Done\n", i)
			}
		}(i)
	}
	fmt.Println("before...")
	go func() {
		time.Sleep(time.Second * 1)
		cancel()
		fmt.Println("...cancel")
		fmt.Printf("%s\n", ctx.Err())
	}()
	fmt.Println("after...")
	time.Sleep(time.Second * 3)
	fmt.Println("end...")
}
func operation1(ctx context.Context) {
	log.Println("operation for id:", ctx.Value(keyID), "completed")

}
func operation2(ctx context.Context) {
	log.Println("operation2 for id:", ctx.Value(keyID), "com")
	select {
	case <-ctx.Done():
		return
	default:
		fmt.Println(ctx.Err())
	}
}

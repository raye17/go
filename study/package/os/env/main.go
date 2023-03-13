package main

import (
	"fmt"
	"os"
)

func main() {
	//for _, e := range os.Environ() {
	//	fmt.Println(e)
	//}
	err := set("name", "name001")
	if err != nil {
		panic(err)
	}
	value := get("name")
	fmt.Println(value)
	func() {
		err := unSet("name")
		if err != nil {
			panic(err)
		}
	}()
	value01 := get("name")
	fmt.Printf("value01:%s", value01)
}
func set(key string, value string) error {
	err := os.Setenv(key, value)
	if err != nil {
		return err
	}
	return nil
}
func unSet(key string) error {
	err := os.Unsetenv(key)
	if err != nil {
		return err
	}
	return nil
}
func get(key string) string {
	value := os.Getenv(key)
	return value
}

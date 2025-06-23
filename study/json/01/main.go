package main

import "encoding/json"

type user struct {
	Name string
	Age  int
	Male bool
}

func main() {
	var u = &user{
		Name: "张三",
		Age:  18,
		Male: true,
	}
	s, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}
	println(string(s))
}

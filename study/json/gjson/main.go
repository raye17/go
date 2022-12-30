package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type res struct {
	Data []Item `json:"data"`
}
type Item struct {
	EName string `json:"eName"`
	Name  string `json:"name"`
	Nid   int    `json:"nid"`
}

func main() {
	var res res
	var resMap map[string]interface{}
	url := "https://studygolang.com/nodes/hot?limit=10?"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("请求出错:", err)
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&res)
	fmt.Printf("result: %+v\n", res)

	json.NewDecoder(resp.Body).Decode(&resMap)
	fmt.Printf("map result: %+v\n", resMap)

}

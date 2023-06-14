package main

import (
	"fmt"
	"net/url"
)

func main() {
	ParseHttp("https://cn.bing.com/search?q=dial&form=CHRDF&sp=-1&sc=10-4&qs=n&cvid=4446907B&ghsh=0")
}
func ParseHttp(path string) {
	u, err := url.Parse(path)
	if err != nil {
		panic(err)
	}
	fmt.Println("u.string", u.String())
	fmt.Println("u:", u)
	fmt.Println("u.rawQuery", u.RawQuery)
	values, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(values)
}

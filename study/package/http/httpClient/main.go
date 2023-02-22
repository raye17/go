package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"study/package/http/httpClient/client"
)

func HiHttp() {
	resp, err := http.Get("https://github.com/raye17/api/v1")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s", string(body[:]))
}
func main() {
	//params := url.Values{}
	//Url, err := url.Parse("https://www.github.com/raye17")
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//params.Set("name", "raye")
	//params.Set("age", "22")
	//Url.RawQuery = params.Encode()
	//urlPath := Url.String()
	//fmt.Println(urlPath)
	//resp, err := http.Get(urlPath)
	//defer resp.Body.Close()
	//body, err := io.ReadAll(resp.Body)
	//fmt.Println(string(body))
	cli := client.Setup(true, false)
	if err := client.DefaultBaidu(); err != nil {
		panic(err)
	}
	if err := client.DoOps(cli); err != nil {
		panic(err)
	}
	c := client.Controller{
		Client: cli,
	}
	if err := c.DoOps(); err != nil {
		panic(err)
	}
	client.Setup(true, true)
	if err := client.DefaultBaidu(); err != nil {
		panic(err)
	}
}

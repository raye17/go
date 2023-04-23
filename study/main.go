package main

import (
	"fmt"
	"net/url"
	"strconv"
)

const name = "sxy"

func main() {
	//dir, _ := os.Getwd()
	//fmt.Println("ll", dir)
	//m := map[string]interface{}{
	//	"map01": map[string]interface{}{
	//		"mmap": "test",
	//	},
	//}
	//fmt.Println(m)
	//re := regexp.MustCompile("(raye){2}")
	//fmt.Println(re)
	//name := name + "raye"
	//fmt.Println(name)
	//s := "https://github.com/raye17/go.git"
	//b := sha256.Sum256([]byte(s))
	//fmt.Println(b)
	//t := hex.EncodeToString(b[:])
	//fmt.Println(t)
}

type student struct {
	name string
	age  int
}

func (s *student) changeName(name string) string {
	s.name = name
	return s.name
}

type func01 func(string) string

func test(s string, func012 func01) {
	fmt.Println(s)
	func012("sss")
}

var entry map[string]charts

type charts []chart

type chart struct {
	url []string
}

func t(baseUrl *url.URL, index map[string]charts) error {
	u := *baseUrl
	i := 0
	for chartName, charts := range index {
		for _, chart := range charts {
			v := url.Values{}
			v.Set("chartName", chartName)
			v.Set("chart test", "versionTest")
			u.RawQuery = v.Encode()
			fmt.Println("RawQuery", u.RawQuery)
			i++
			c := fmt.Sprintf(u.String() + "id:" + strconv.Itoa(i))
			fmt.Println("c", c)
			chart.url = []string{c}
		}
	}
	return nil
}

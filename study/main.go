package main

import (
	"fmt"
	"net/url"
	"strconv"
)

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

func main() {
	//var chart01 = chart{
	//	url: []string{"http://local:1111/?abc=123&fgh=456", "http://local:2222/test"},
	//}
	//var chart002 = chart{
	//	url: []string{"https://localhost:3333/test033/fg", "ch02"},
	//}
	//var chs = charts{chart01, chart002}
	//entry = make(map[string]charts)
	//entry["entry01"] = chs
	//entry["entry02"] = charts{
	//	chart{
	//		url: []string{"entry-url01", "entry-url02"},
	//	},
	//	chart{
	//		url: []string{"entry02-url001"},
	//	},
	//}
	//u := url.URL{}
	//t(&u, entry)
	//fmt.Println("after t ...")
	//fmt.Println(entry)
	//fmt.Println(u)
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

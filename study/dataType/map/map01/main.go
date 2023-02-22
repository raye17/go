package main

import "fmt"

func main() {
	//元素类型为map的切片
	var mapSlice = make([]map[string]int, 8)
	for i := 0; i < 8; i++ {
		mapSlice[i] = make(map[string]int, 8)
	}
	mapSlice[0]["raye"] = 19
	mapSlice[1]["lcx"] = 18
	mapSlice[1]["sxy"] = 199
	mapSlice[7]["xxx"] = 100
	for k := range mapSlice {
		for i, v := range mapSlice[k] {
			fmt.Println(k, i, v)
		}
	}
	//值为map的切片
	fmt.Println("值为map的切片", "\n", "\n", "***")
	var sliceMap = make(map[string][]string, 8)
	_, ok := sliceMap["中国"]
	if ok {
		fmt.Println(sliceMap["中国"])
	} else {
		sliceMap["中国"] = make([]string, 8)
		sliceMap["河南"] = make([]string, 8)
		sliceMap["周口"] = make([]string, 8)
		sliceMap["苏州"] = make([]string, 8)
		sliceMap["杭州"] = make([]string, 8)
		sliceMap["江苏"] = make([]string, 8)
		sliceMap["中国"][0] = "china"
		sliceMap["中国"][1] = "China CN"
		sliceMap["河南"][2] = "郑州"
		sliceMap["河南"][4] = "商丘"
		sliceMap["江苏"][1] = "苏州"
		sliceMap["苏州"][3] = "虎丘"
		sliceMap["苏州"][2] = "吴中"
		sliceMap["杭州"][3] = "西湖"
		sliceMap["周口"][4] = "郸城"
	}
	fmt.Println(sliceMap)
}

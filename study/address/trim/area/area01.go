package area

import (
	"encoding/json"
	"strings"
)

type AreaTree struct {
	AreaInfo
	Children []AreaTree
}
type AreaInfo struct {
	//bankQuery.Data
	//区域信息
	Id        int    `json:"id"`         //id
	Pid       int    `json:"pid"`        //父级id
	Name      string `json:"name"`       //地区名称
	Level     int    `json:"level"`      //层级
	AreaCode  string `json:"area_code"`  //区域代码
	TownCode  string `json:"town_code"`  //乡镇代码
	CountryId int    `json:"country_id"` //国家id
}

var chinaAreaTree []AreaTree

func GetChinaAreaTree() (data []AreaTree, err error) {
	if chinaAreaTree != nil {
		return chinaAreaTree[0].Children, nil
	}
	chinaAreaTree, err = GetAreaTree()
	if err != nil {
		return
	}
	return chinaAreaTree[0].Children, nil
}
func GetAreaTree() (data []AreaTree, err error) {
	var listData []AreaInfo
	err = json.Unmarshal([]byte(areaList), &listData)
	if err != nil {
		return
	}
	for _, v := range listData {
		if v.Pid == 0 {
			v := v
			var newArea = AreaTree{
				AreaInfo: v,
			}
			newArea.Children = buildChildrenTree(listData, newArea.Id, newArea.Name)
			data = append(data, newArea)
			break
		}
	}
	return
}
func buildChildrenTree(data []AreaInfo, pid int, pName string) (children []AreaTree) {
	for _, v := range data {
		if v.Pid == pid {
			v := v
			var newArea = AreaTree{
				AreaInfo: v,
			}
			//市辖区的名称改为省份名称
			if newArea.Name == "市辖区" {
				newArea.Name = pName
			}
			newArea.Children = buildChildrenTree(data, newArea.Id, newArea.Name)
			children = append(children, newArea)
		}
	}
	return
}
func AddressParseToFourLevel(addr string, inputDeep ...int) (province string, city string, areas string, address string) {
	deep := 1
	if inputDeep != nil {
		deep = inputDeep[0]
	}
	addrRunes := []rune(addr)
	found := 0
	var foundArea AreaTree
	areaTree, err := GetChinaAreaTree()
	if err != nil {
		return "", "", "", addr
	}
	var tmp []rune
	for _, v := range addrRunes {
		if string(v) == "" || string(v) == " " || string(v) == "," || string(v) == "，" {
			continue
		}
		tmp = append(tmp, v)
		if len(tmp) <= 2 || found > 3 {
			continue
		}
		switch found {
		case 0: //找省
			for _, areaItem := range areaTree {
				//fmt.Println(string(tmp) + " compare " + areaItem.Name)
				if areaItem.Name == string(tmp) {
					province = string(tmp)
					tmp = []rune{}
					found++
					foundArea = areaItem
					break
				}
			}
		case 1: //找市
			for _, areaItem := range foundArea.Children {
				if areaItem.Name == string(tmp) {
					city = string(tmp)
					tmp = []rune{}
					found++
					foundArea = areaItem
					break
				}
			}
		case 2: //找区
			for _, areaItem := range foundArea.Children {
				if areaItem.Name == string(tmp) {
					areas = string(tmp)
					tmp = []rune{}
					found++
					foundArea = areaItem
					break
				}
			}
		}
	}
	deep++
	if deep > 5 {
		address = string(tmp)
	} else {
		if strings.Contains(string(tmp), "省") && strings.Contains(string(tmp), "市") {
			return AddressParseToFourLevel(string(tmp), deep)
		}
		address = string(tmp)
	}
	return
}

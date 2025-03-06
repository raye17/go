package string

import (
	"os"
	"strconv"
	"strings"
)

// 格式化输出
func FloToStr(num float64) (res []string) {
	res = append(res, strconv.FormatFloat(num, 'f', -1, 64))
	res = append(res, strconv.FormatFloat(num, 'f', 0, 64))
	res = append(res, strconv.FormatFloat(num, 'f', 2, 64))
	return
}

// add
func Add(a, b int) (c int, err error) {
	file, err := os.Create("./res.md")
	if err != nil {
		return c, err
	}
	c = a + b
	file.WriteString(strconv.Itoa(c))
	return
}
func Sfrim(str string) []string {
	res := strings.Split(str, ";")
	return res
}

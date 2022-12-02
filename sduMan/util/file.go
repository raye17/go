package util

import (
	"encoding/csv"
	"fmt"
	"os"
	"raye/stu.com/manage"
	"reflect"
)

func StructToSlice(stuMgs manage.Student) []string {
	v := reflect.ValueOf(stuMgs)

	ss := make([]string, v.NumField())
	for i := range ss {
		ss[i] = fmt.Sprintf("%v\t", v.Field(i))
	}
	return ss
}

func EditFile(stuMgs manage.Student) {
	file, err := os.Open("./data/stuInfo.md")
	defer file.Close()
	if err != nil && os.IsNotExist(err) {
		file, err := os.Create("./data/stuInfo.md")
		Check(err)
		defer file.Close()
		//写入字段标题
		w := csv.NewWriter(file)
		title := []string{"id\tname\tage\tgender\tclass"}
		w.Write(title)
		stuInfo := StructToSlice(stuMgs)
		w.Write(stuInfo)
		w.Flush()
	} else {
		txt, err := os.OpenFile("./data/stuInfo.md", os.O_APPEND|os.O_RDWR, 0666)
		defer txt.Close()
		Check(err)
		w := csv.NewWriter(txt)
		stuInfo := StructToSlice(stuMgs)
		w.Write(stuInfo)
		w.Flush()
	}
}

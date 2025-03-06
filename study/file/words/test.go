package words

import (
	"fmt"
	"time"

	docx "github.com/lukasjarosch/go-docx"
)

func ReadDocxFile() {
	file := "./doc/test01.docx"
	doc, err := docx.Open(file)
	if err != nil {
		panic(err)
	}
	replaceMap := docx.PlaceholderMap{
		"name": "孙肖扬",
		"age":  19,
	}
	if err = doc.ReplaceAll(replaceMap); err != nil {
		panic(err)
	}
	now := time.Now().UnixMilli()
	tmpPath := fmt.Sprintf("./data/%d.docx", now)
	err = doc.WriteToFile(tmpPath)
	if err != nil {
		panic(err)
	}
}

package utils

import (
	"fmt"
	"os"
	"ssh/demo/model"
	"strconv"
	"strings"
)

func ConvertFieldToInt(field interface{}, fieldName string) (int, bool, error) {
	if field == nil {
		return 0, false, nil
	}

	strVal, ok := field.(string)
	if !ok {
		return 0, false, fmt.Errorf("%s字段不是字符串类型", fieldName)
	}

	if strVal == "" {
		return 0, false, nil
	}

	intVal, err := strconv.Atoi(strVal)
	if err != nil {
		return 0, false, fmt.Errorf("转换%s失败: %v", fieldName, err)
	}

	return intVal, true, nil
}

// 修改文件写入部分
func WriteRecordsToFile(records []model.ResultRecord, filename string) error {
	var lines []string
	for _, r := range records {
		lines = append(lines, fmt.Sprintf("%s\t%s\t%s", r.UUID, r.Message, r.OperateDate))
	}

	// 检查文件是否存在
	if _, err := os.Stat(filename); err == nil {
		// 文件存在，以追加模式打开
		file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer file.Close()

		// 写入换行符和新增内容
		if _, err = file.WriteString("\n" + strings.Join(lines, "\n")); err != nil {
			return err
		}
	} else {
		// 文件不存在，创建新文件
		return os.WriteFile(filename, []byte(strings.Join(lines, "\n")), 0644)
	}
	return nil
}

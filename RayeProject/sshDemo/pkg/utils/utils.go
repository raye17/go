package utils

import (
	"fmt"
	"strconv"
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

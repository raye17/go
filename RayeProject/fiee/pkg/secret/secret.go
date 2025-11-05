package secret

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"test001/pkg/secret/aes"
)

func GetPositionCode(positionName string) string {
	code := "other"
	positionMap := map[string]string{"宣传部": "xcb", "鉴证科": "jzk", "经纪人": "jjr", "普通用户": "com"}

	if v, ok := positionMap[positionName]; ok {
		code = v
	}

	return code
}

func CombineSecret(positionID, departmentID, toke string) (string, error) {
	abc := positionID + "(~!@)" + departmentID + "(~!@)" + toke
	b, err := aes.AesEcrypt([]byte(abc), aes.PwdKey)
	if err != nil {
		return "", errors.New("解析错误")
	}

	return hex.EncodeToString(b), nil
}

func getToken(token string) (string, string, string, error) {

	if strings.Contains(token, "(~!@)") == false {
		return "", "", "", errors.New("解析错误")
	}

	str1 := strings.Split(token, "(~!@)")

	if len(str1) != 3 {
		return "", "", "", errors.New("解析数量错误")
	}

	return str1[2], str1[1], str1[0], nil
}
func getTokenV2(token string) (string, error) {

	if strings.Contains(token, "(~!@)") == false {
		return "", errors.New("解析错误")
	}

	str1 := strings.Split(token, "(~!@)")

	if len(str1) != 3 {
		return "", errors.New("解析数量错误")
	}

	return str1[2], nil
}

func GetJwtFromStr(authorization string) (string, uint64, uint64, error) {
	var departmentID, positionID uint64
	departmentID = 0
	positionID = 0

	tokenByte, err := hex.DecodeString(authorization)
	if err != nil {
		return "", departmentID, positionID, err
	}
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println("panic信息:", msg, "---recover恢复---")
		}
	}()

	token, err := aes.AesDeCrypt(tokenByte, aes.PwdKey)

	if err != nil {
		return "", departmentID, positionID, err
	}

	//解密下
	jwt, departmentIDStr, positionIDStr, err := getToken(string(token))
	if err != nil {
		return "", departmentID, positionID, err
	}

	departmentID, err = strconv.ParseUint(departmentIDStr, 10, 64)
	if err != nil {
		return "", departmentID, positionID, err
	}

	positionID, err = strconv.ParseUint(positionIDStr, 10, 64)
	if err != nil {
		return "", departmentID, positionID, err
	}

	return jwt, departmentID, positionID, nil

}

func GetJwtFromStrV2(authorization string) (string, error) {

	tokenByte, err := hex.DecodeString(authorization)
	if err != nil {
		return "", err
	}
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println("panic信息:", msg, "---recover恢复---")
		}
	}()

	token, err := aes.AesDeCrypt(tokenByte, aes.PwdKey)

	if err != nil {
		return "", err
	}

	//解密下
	jwt, err := getTokenV2(string(token))

	if err != nil {
		return "", err
	}

	return jwt, nil

}

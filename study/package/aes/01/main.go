package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	authorization := "df6498d8162138dd1e67401276d0d7f7ea200a2e29b4326bf2346ded05c3742647cded1d754e6809b54754a349304399962065e04040e502227f819815cfd5d1edd69c62473b8eb6fff69b26d5619b6e073becd826445e1f3415d2e779710150246d4f4601d749690a080fe51265c2b1c6c120393573713eab73888d36b796733331b4f5f495ceebedb1089a7a0df5de6c4bf18aa16816a0b713d0321609ccbe840541af2ce3b99a40b224d9631521a9ea46f07abdaae4f167397bfaf9d0ca90476f1568713f5255afd7d5cbb09e3dae052af09c4302889225660366a1ffed6d68f4f483222815384f9ff72d8daf8a72508e8c12891356b6fb9e3ed8201dc94f43ed98e3d4037fbe7681c785bb61b907fa8ad5ce1370e5a907f3173c388a1d6be3a49050b7566e2002834a61b44502fc"
	jwt, nowDepartmentID, id, err := GetJwtFromStr(authorization)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("jwt:", jwt)
	fmt.Println("nowDepartmentID:", nowDepartmentID)
	fmt.Println("id:", id)
}

var PwdKey = []byte("tyfon918tyfon918")

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

	token, err := AesDeCrypt(tokenByte, PwdKey)

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

// 实现解密
func AesDeCrypt(cypted []byte, key []byte) ([]byte, error) {
	//创建加密算法实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//获取块大小
	blockSize := block.BlockSize()
	//创建加密客户端实例
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(cypted))
	//这个函数也可以用来解密
	blockMode.CryptBlocks(origData, cypted)
	//去除填充字符串
	origData, err = PKCS7UnPadding(origData)
	if err != nil {
		return nil, err
	}
	return origData, err
}

// 填充的反向操作，删除填充字符串
func PKCS7UnPadding(origData []byte) ([]byte, error) {
	//获取数据长度
	length := len(origData)
	if length == 0 {
		return nil, errors.New("加密字符串错误！")
	} else {
		//获取填充字符串长度
		unpadding := int(origData[length-1])
		//截取切片，删除填充字节，并且返回明文
		return origData[:(length - unpadding)], nil
	}
}

package aes

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
)

type CommonSecret struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type PayInfo struct {
	Id      uint   `json:"id"`
	OrderNo string `json:"orderNo"`
	Uuid    string `json:"uuid"`
}

type AesSetting struct {
	Pwd []byte
}

func CreateSecretDemoPay(id uint, orderNo, uuid string) (string, error) {

	aesSetting := AesSetting{Pwd: []byte("tyfon918tyfon918")}

	req := PayInfo{
		Id:      id,
		OrderNo: orderNo,
		Uuid:    uuid,
	}

	return aesSetting.CreateSecretPay(req)
}

func EncryptionSecretPay(token string) (*PayInfo, error) {

	aesSetting := AesSetting{Pwd: []byte("tyfon918tyfon918")}

	return aesSetting.EncryptionSecretPay(token)
}

func CreateSecretDemo(id uint, name string) (string, error) {

	aesSetting := AesSetting{Pwd: []byte("tyfon918tyfon918")}

	req := CommonSecret{
		Id:   id,
		Name: name,
	}

	return aesSetting.CreateSecret(req)
}

func EncryptionSecret(token string) (*CommonSecret, error) {

	aesSetting := AesSetting{Pwd: []byte("tyfon918tyfon918")}

	return aesSetting.EncryptionSecret(token)
}

// CreateSecret 加密
func (a *AesSetting) CreateSecret(userMsg CommonSecret) (string, error) {
	v, err := json.Marshal(userMsg)
	if err != nil {
		return "", err
	}

	b, err := AesEcrypt(v, a.Pwd)
	if err != nil {
		return "", errors.New("解析错误")
	}

	return hex.EncodeToString(b), nil
}

// EncryptionSecret 解密
func (a *AesSetting) EncryptionSecret(token string) (userMsg *CommonSecret, err error) {
	tokenByte, err := hex.DecodeString(token)
	if err != nil {
		return nil, err
	}
	fmt.Println(1)
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println("panic信息:", msg, "---recover恢复---")
			err = errors.New("token 解析失败")
		}
	}()

	fmt.Println(2)
	res, err := AesDeCrypt(tokenByte, a.Pwd)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &userMsg)

	if err != nil {
		return nil, err
	}

	return userMsg, err
}

// CreateSecret 加密
func (a *AesSetting) CreateSecretPay(userMsg PayInfo) (string, error) {
	v, err := json.Marshal(userMsg)
	if err != nil {
		return "", err
	}

	b, err := AesEcrypt(v, a.Pwd)
	if err != nil {
		return "", errors.New("解析错误")
	}

	return hex.EncodeToString(b), nil
}

// EncryptionSecret 解密
func (a *AesSetting) EncryptionSecretPay(token string) (userMsg *PayInfo, err error) {
	tokenByte, err := hex.DecodeString(token)
	if err != nil {
		return nil, err
	}
	fmt.Println(1)
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println("panic信息:", msg, "---recover恢复---")
			err = errors.New("token 解析失败")
		}
	}()

	fmt.Println(2)
	res, err := AesDeCrypt(tokenByte, a.Pwd)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &userMsg)

	if err != nil {
		return nil, err
	}

	return userMsg, err
}

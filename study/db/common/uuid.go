package common

import "github.com/gofrs/uuid"

// 生成uuid
func NewUuid() (string, error) {
	// 生成UUID
	uuid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return uuid.String(), nil
}

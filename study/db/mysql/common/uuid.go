package common

import "github.com/google/uuid"

// NewUuid 生成uuid
func NewUuid() (string, error) {
	// 生成UUID
	uuid, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return uuid.String(), nil
}

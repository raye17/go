package common

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func GenerateFilingKey(filingDate, formType, finalLink string) string {
	data := fmt.Sprintf("%s_%s_%s", filingDate, formType, finalLink)
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}

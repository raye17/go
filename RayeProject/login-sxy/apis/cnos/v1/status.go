package v1

import (
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
)

func (u *UserStatus) Bytes() ([]byte, error) {
	bytes, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}
	newStr := fmt.Sprintf(`"{status":%s}`, string(bytes))
	glog.V(5).Infof("status body", newStr)
	return []byte(newStr), nil
}

package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"test001/api/position"
	"test001/pkg/e"

	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"github.com/gin-gonic/gin"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

// Response 基础序列化器
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Code   int         `json:"code"`
	Error  error       `json:"error"`
	Err    string      `json:"err"`
	Keys   []string    `json:"keys"`
}

func Success(c *gin.Context, data proto.Message) {
	m := jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  true,
		EmitDefaults: false,
		Indent:       "",
		AnyResolver:  nil,
	}

	dataMap := make(map[string]interface{})

	fmt.Printf(data.String())
	str, err := m.MarshalToString(data)
	fmt.Printf(str)
	if err != nil {

	}
	err = json.Unmarshal([]byte(str), &dataMap)
	if err != nil {

	}

	c.JSON(http.StatusOK, Response{
		Status: e.Ok,
		Code:   e.SUCCESS,
		Data:   dataMap,
	})

	c.Abort()
}

// Success Success(c,someMap)
func SuccessV2(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Status: e.Ok,
		Code:   e.SUCCESS,
		Data:   data,
	})
	c.Abort()
}

// Error 统一错误返回
func Error(c *gin.Context, code int, err error, msg ...string) {
	logger.Error(err)
	status := e.Failed
	if code == e.NotLogin {
		status = e.NotLogin
	}
	c.JSON(e.Success, Response{
		Status: status,
		Msg:    err.Error(),
		Data:   nil,
	})

	c.Abort()
}

func IsHaveAuth(c *gin.Context, ID uint64, key string, departmentId uint64) bool {
	isLeader := false

	req := position.DoIHavaAuthRequest{
		UserId: ID,
		//Url:          e.SellerBossKey,
		Url:          key,
		DepartmentID: departmentId,
	}

	res, err := PositionProvider.DoIHavaAuth(c, &req)

	if err == nil && res.Hava == true {
		isLeader = true
	}

	return isLeader
}

// 废弃

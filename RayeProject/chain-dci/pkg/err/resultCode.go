package err

var ResultCode = map[string]string{
	"OK":               "正常返回",
	"BAD_REQUEST":      "请求参数错误",
	"PERMISSION_ERROR": "权限错误",
	"BUSINESS_ERROR":   "业务内容错误",
	"SERVER_ERROR":     "服务器错误",
}

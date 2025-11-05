package e

const (
	SERVER_CONFIG         = "conf.ini"
	SERVER_DUBBOGO_CONFIG = "dubbogo.yaml"
	MODE_ENV              = "MODE_ENV"
)

var MsgFlags = map[int]string{
	SUCCESS:               "操作成功",
	UpdatePasswordSuccess: "修改密码成功",
	NotExistInentifier:    "该第三方账号未绑定",
	ERROR:                 "fail",
	InvalidParams:         "请求参数错误",
	BindError:             "参数绑定错误，类型不一致",
	JsonUnmarshal:         "Json解析错误",

	ErrorExistNick:          "已存在该昵称",
	ErrorExistUser:          "已存在该用户名",
	ErrorNotExistUser:       "该用户不存在",
	ErrorNotCompare:         "账号密码错误",
	ErrorNotComparePassword: "两次密码输入不一致",
	ErrorFailEncryption:     "加密失败",
	ErrorNotExistProduct:    "该商品不存在",
	ErrorNotExistAddress:    "该收获地址不存在",
	ErrorExistFavorite:      "已收藏该商品",
	ErrorGetUserInfo:        "获取用户信息错误",
	ErrorGetDepart:          "获取部门信息错误",
	ErrorUpdateAw:           "同步画作信息错误",

	ErrorBossCheckTokenFail:        "商家的Token鉴权失败",
	ErrorBossCheckTokenTimeout:     "商家TOken已超时",
	ErrorBossToken:                 "商家的Token生成失败",
	ErrorBoss:                      "商家Token错误",
	ErrorBossInsufficientAuthority: "商家权限不足",
	ErrorBossProduct:               "商家读文件错误",

	ErrorAuthCheckTokenFail:        "Token鉴权失败",
	ErrorAuthCheckTokenTimeout:     "TOken已超时",
	ErrorAuthToken:                 "Token生成失败",
	ErrorAuth:                      "Token错误",
	ErrorAuthInsufficientAuthority: "权限不足",
	ErrorReadFile:                  "读文件失败",
	ErrorSendEmail:                 "发送邮件失败",
	ErrorCallApi:                   "调用接口失败",
	ErrorUnmarshalJson:             "解码JSON失败",

	ErrorUploadFile:    "上传失败",
	ErrorAdminFindUser: "管理员查询用户失败",

	ErrorDatabase: "数据库操作出错,请重试",

	ErrorOss: "OSS配置错误",

	ErrorExistShopName:    "店铺已被注册，请检查店铺名称和统一社会信用码",
	ErrorNotExistShopName: "店铺不存在",
	ErrorNotAdmin:         "非管理员",

	InvalidToken: "Token验证失败",

	ErrorUploadVideoCover: "视频截取封面错误",
	ErrorUploadValidParam: "上传参数非法",
	ErrorFileReadErr:      "读取文件错误",
	ErrorFileNotExists:    "文件不存在",
	ErrorChunkNotGt:       "分块数量不一致",
	ErrorChunk:            "读取分块错误",
	ErrorUploadBos:        "上传bos错误",
	ErrorFileCreate:       "文件创建错误",
	ERROR_UID:             "uid创建错误",

	ErrNoDomain: "环境变量必须要有",
	ErrTelNum:   "手机号码错误",
	ErrNoCode:   "验证码必须要有",
	ErrNoID:     "ID缺少",
	ErrNickName: "昵称长度超过了20个字符或者缺失",
	InvalidID:   "身份证长度18位",
	InvalidPas:  "密码不小于6位",

	ErrStatus:        "状态非法",
	ErrNoType:        "缺少类型",
	ErrNoUserID:      "缺少用户ID",
	ErrNoName:        "缺少名称",
	ErrNoDepCode:     "缺少部门code",
	ErrNoTitle:       "缺少标题",
	ErrNoUrl:         "缺少url",
	ErrNoMethod:      "缺少method",
	ErrNotDep:        "缺少部门",
	ErrCreateQr:      "生成二维码错误",
	ErrNotSellerBoss: "当前人员身份非销售总监,不能操作",
	ErrWrongDate:     "时间数据填写错误",
	ErrLoginSeller:   "您暂时没有权限登录销售宝，请联系管理员",
}

// GetMsg 获取状态码对应信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}

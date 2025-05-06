package msg

const (
	ErrFileIsEmpty              = "未获取到相关文件"
	ErrCreateClientToken        = "创建clientToken失败"
	ErrAddDciUser               = "著作权人申领DCI信息失败"
	ErrUpdateDciUser            = "著作权人更新DCI信息失败"
	ErrQueryDciUser             = "著作权人查询DCI信息失败"
	ErrCreateDciPreregistration = "作品申领DCI失败"
	ErrQueryDciPreregistration  = "作品查询DCI失败"
	ErrCreateDciRegistration    = "发起数登申请失败"
	ErrQueryDciRegistration     = "查询数登申请失败"
	ErrGetDciPayurl             = "数登支付链接获取失败"
	ErrQueryDciPay              = "数登支付查询失败"
	ErrGetDciRegistrationcert   = "获取数登证书下载失败"
	ErrRetryDciRegistration     = "补正申请失败"
	ErrCloseDciRegistration     = "数登停止申请失败"
	ErrSubmitDciFeedback        = "作品DCI申诉失败"
)

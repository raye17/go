package err

var DciDigitalRegStatusEnum = map[string]string{
	"CHECK":             "初审中",
	"CHECK_FAIL":        "初审不通过",
	"PAY":               "待支付",
	"REVIEW":            "复审中",
	"TOBE_AMEND":        "待补正",
	"AMEND_CHECK_FAIL":  "待补正审核失败",
	"FINISH":            "数登完成",
	"DISREGARD":         "不予处理",
	"REGISTRATION_STOP": "申请停止",
}

package err

var WorkStatus = map[string]string{
	"DCI_PRE_REG_AUDITING":     "申领审核中",
	"DCI_PRE_REG_FINISH":       "申领成功",
	"DCI_PRE_REG_FAIL":         "申领失败",
	"DCI_PRE_REG_CANCEL":       "dci撤销",
	"DCI_PRE_REG_TOBE_CONFIRM": "申领待确认，账户余额不足",
}

package logic

import (
	"bytes"
	"chain-dci/pkg/msg"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"go.uber.org/zap"
)

type IDci interface {
	// DciAddUser	注册DCI账号   每位著作权人 都需要注册此账号
	// DciUpdateUser	更新DCI用户信息   著作权人 证件信息更新
	// DciQueryUser	查询DCI用户信息    著作权人 信息查询

	// DciPreregistration	DCI申领   申请版权的作品 需要 先进行 DCI申领 获取到 DCI作品ID
	// DciQueryPreregistration	查询dci申领信息   通过 DCI作品ID   获取到 DCI申领的详细信息
	// DciQueryPreregPublication	查询dci申领公示地址

	// DciRegistration 发起数登申请  在发起数字版权登记（简称：数登）申请前，请先进行DCI申领，只有申领完成的DCI才能发起数登申请
	// DciQueryRegistration 查询数登申请  用户可通过主动查询的方式获取数登结果，支持按数登id或发起数登的DCI申领id来查询
	// DciGetPayUrl 获取数登支付链接  初审通过后，用户通过主动查询的方式获取数登支付链接，该接口在北京时间每日06:00-23:00可用，支付链接有效期15分钟，若支付链接过期可再次调用接口刷新支付链接
	// DciQueryPay 数登支付查询  用户完成支付后，通过数登支付查询接口获取支付结果
	// DciGetRegistrationCert 获取数登证书下载链接  对于申请成功的数登，可通过该接口获取数登证书下载链接，数登证书下载次数限制为10次/个/月，下载链接有效期为15天
	// RetryDigitalRegister 补正申请  对于初审不通过、待补证和待补证审核不通过的数登，请按照反馈原因调整数登申请信息，完成调整后通过补正申请接口对数登进行修改
	// CloseDigitalRegister 数登停止申请  对于待补证和待补证审核不通过的数登，如不想进行补正，请调用数登申请停止接口，告知系统申请停止
}

type Dci struct {
}

func createToken(timestamp int64, keyword ...string) (token string, err error) {
	var b bytes.Buffer
	for i := 0; i < len(keyword); i++ {
		b.WriteString(keyword[i])
	}
	b.WriteString(fmt.Sprint(timestamp))
	h := sha256.New()

	if _, err = h.Write([]byte(b.String())); err != nil {
		zap.L().Error("createToken sha256 err", zap.Error(err))
		err = errors.New(msg.ErrorSha256Write)
		return
	}
	s := h.Sum(nil)
	token = hex.EncodeToString(s)
	return
}

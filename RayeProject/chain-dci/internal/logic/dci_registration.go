package logic

import (
	"chain-dci/pb/dci"
	"chain-dci/pkg/app"
	errCommon "chain-dci/pkg/err"
	"chain-dci/pkg/msg"
	"fmt"
	bccrClient "github.com/antchain-openapi-sdk-go/bccr/client"
	"github.com/jinzhu/copier"
	"time"
)

type IDciRegistration interface {
	CreateDciRegistration(req *dci.CreateDciRegistrationRequest) (res *dci.CreateDciRegistrationResponse, err error)
	QueryDciRegistration(req *dci.QueryDciRegistrationRequest) (res *dci.QueryDciRegistrationResponse, err error)
	GetDciPayUrl(req *dci.GetDciPayUrlRequest) (res *dci.GetDciPayUrlResponse, err error)
	QueryDciPay(req *dci.QueryDciPayRequest) (res *dci.QueryDciPayResponse, err error)
	GetDciRegistrationcert(req *dci.GetDciRegistrationcertRequest) (res *dci.GetDciRegistrationcertResponse, err error)
	RetryDciRegistration(req *dci.RetryDciRegistrationRequest) (res *dci.RetryDciRegistrationResponse, err error)
	CloseDciRegistration(req *dci.CloseDciRegistrationRequest) (res *dci.CloseDciRegistrationResponse, err error)
}

type DciRegistration struct {
}

// CreateDciRegistration 发起数登申请
func (r *DciRegistration) CreateDciRegistration(req *dci.CreateDciRegistrationRequest) (res *dci.CreateDciRegistrationResponse, err error) {
	errCommon.NoReturnInfo(req, "发起数登申请 参数信息: ")

	res = new(dci.CreateDciRegistrationResponse)

	createDciRegistrationRequest := new(bccrClient.CreateDciRegistrationRequest)
	_ = copier.CopyWithOption(&createDciRegistrationRequest, req, copier.Option{DeepCopy: false})
	fmt.Printf("req : %+v\n", req)
	fmt.Println("======================================================= ============== ===========00")
	fmt.Printf("createDciRegistrationRequest : %+v\n", createDciRegistrationRequest)

	clientToken, err := createToken(time.Now().UnixMilli(), req.DciContentId, app.ModuleClients.SfNode.Generate().Base64())
	if err != nil {
		return nil, errCommon.ReturnError(err, msg.ErrCreateClientToken, "创建clientToken 失败: ")
	}
	createDciRegistrationRequest.SetClientToken(clientToken)

	createDciPreregistrationResponse, err := app.ModuleClients.BccrClient.CreateDciRegistration(createDciRegistrationRequest)
	if err != nil {
		return nil, errCommon.ReturnError(err, msg.ErrCreateDciRegistration, "发起数登申请 失败: ")
	}

	errCommon.NoReturnInfo(createDciPreregistrationResponse, "发起数登申请 成功: ")

	_ = copier.CopyWithOption(&res, createDciPreregistrationResponse, copier.Option{DeepCopy: false})

	return
}

// QueryDciRegistration 查询数登申请
func (r *DciRegistration) QueryDciRegistration(req *dci.QueryDciRegistrationRequest) (res *dci.QueryDciRegistrationResponse, err error) {
	errCommon.NoReturnInfo(req, "查询数登申请 参数信息: ")

	res = new(dci.QueryDciRegistrationResponse)

	queryDciRegistrationRequest := new(bccrClient.QueryDciRegistrationRequest)
	_ = copier.CopyWithOption(&queryDciRegistrationRequest, req, copier.Option{DeepCopy: false})

	queryDciRegistrationResponse, err := app.ModuleClients.BccrClient.QueryDciRegistration(queryDciRegistrationRequest)
	if err != nil {
		return nil, errCommon.ReturnError(err, msg.ErrQueryDciRegistration, "查询数登申请 失败: ")
	}

	errCommon.NoReturnInfo(queryDciRegistrationResponse, "查询数登申请 成功: ")

	_ = copier.CopyWithOption(&res, queryDciRegistrationResponse, copier.Option{DeepCopy: false})

	return
}

// GetDciPayUrl 数登支付链接获取
func (r *DciRegistration) GetDciPayUrl(req *dci.GetDciPayUrlRequest) (res *dci.GetDciPayUrlResponse, err error) {
	errCommon.NoReturnInfo(req, "数登支付链接获取 参数信息: ")

	res = new(dci.GetDciPayUrlResponse)

	getDciPayUrlRequest := new(bccrClient.GetDciPayurlRequest)
	_ = copier.CopyWithOption(&getDciPayUrlRequest, req, copier.Option{DeepCopy: false})

	clientToken, err := createToken(time.Now().UnixMilli(), req.DigitalRegisterId, app.ModuleClients.SfNode.Generate().Base64())
	if err != nil {
		return nil, errCommon.ReturnError(err, msg.ErrCreateClientToken, "创建clientToken 失败: ")
	}
	getDciPayUrlRequest.SetClientToken(clientToken)

	getDciPayUrlResponse, err := app.ModuleClients.BccrClient.GetDciPayurl(getDciPayUrlRequest)
	if err != nil {
		return nil, errCommon.ReturnError(err, msg.ErrGetDciPayurl, "数登支付链接获取 失败: ")
	}

	errCommon.NoReturnInfo(getDciPayUrlResponse, "数登支付链接获取 成功: ")

	_ = copier.CopyWithOption(&res, getDciPayUrlResponse, copier.Option{DeepCopy: false})

	// 获取支付金额
	if res.PayUrl != "" {
		res.Amount = getPayUrlAmount(res.PayUrl)
	}

	return
}

// QueryDciPay 数登支付查询
func (r *DciRegistration) QueryDciPay(req *dci.QueryDciPayRequest) (res *dci.QueryDciPayResponse, err error) {
	errCommon.NoReturnInfo(req, "数登支付查询 参数信息: ")

	res = new(dci.QueryDciPayResponse)

	queryDciPayRequest := new(bccrClient.QueryDciPayRequest)
	_ = copier.CopyWithOption(&queryDciPayRequest, req, copier.Option{DeepCopy: false})

	queryDciPayResponse, err := app.ModuleClients.BccrClient.QueryDciPay(queryDciPayRequest)
	if err != nil {
		return nil, errCommon.ReturnError(err, msg.ErrQueryDciPay, "数登支付查询 失败: ")
	}

	errCommon.NoReturnInfo(queryDciPayResponse, "数登支付查询 成功: ")

	_ = copier.CopyWithOption(&res, queryDciPayResponse, copier.Option{DeepCopy: false})

	return
}

// GetDciRegistrationcert 获取数登证书下载
func (r *DciRegistration) GetDciRegistrationcert(req *dci.GetDciRegistrationcertRequest) (res *dci.GetDciRegistrationcertResponse, err error) {
	errCommon.NoReturnInfo(req, "获取数登证书下载 参数信息: ")

	res = new(dci.GetDciRegistrationcertResponse)

	getDciRegistrationcertRequest := new(bccrClient.GetDciRegistrationcertRequest)
	_ = copier.CopyWithOption(&getDciRegistrationcertRequest, req, copier.Option{DeepCopy: false})

	clientToken, err := createToken(time.Now().UnixMilli(), req.DigitalRegisterId, app.ModuleClients.SfNode.Generate().Base64())
	if err != nil {
		return nil, errCommon.ReturnError(err, msg.ErrCreateClientToken, "创建clientToken 失败: ")
	}
	getDciRegistrationcertRequest.SetClientToken(clientToken)

	getDciRegistrationcertResponse, err := app.ModuleClients.BccrClient.GetDciRegistrationcert(getDciRegistrationcertRequest)
	if err != nil {
		return nil, errCommon.ReturnError(err, msg.ErrGetDciRegistrationcert, "获取数登证书下载 失败: ")
	}

	errCommon.NoReturnInfo(getDciRegistrationcertResponse, "获取数登证书下载 成功: ")

	_ = copier.CopyWithOption(&res, getDciRegistrationcertResponse, copier.Option{DeepCopy: false})

	return
}

// RetryDciRegistration 补正申请
func (r *DciRegistration) RetryDciRegistration(req *dci.RetryDciRegistrationRequest) (res *dci.RetryDciRegistrationResponse, err error) {
	errCommon.NoReturnInfo(req, "补正申请 参数信息: ")

	res = new(dci.RetryDciRegistrationResponse)

	retryDciRegistrationRequest := new(bccrClient.RetryDciRegistrationRequest)
	_ = copier.CopyWithOption(&retryDciRegistrationRequest, req, copier.Option{DeepCopy: false})

	clientToken, err := createToken(time.Now().UnixMilli(), req.DigitalRegisterId, req.DciContentId, app.ModuleClients.SfNode.Generate().Base64())
	if err != nil {
		return nil, errCommon.ReturnError(err, msg.ErrCreateClientToken, "创建clientToken 失败: ")
	}
	retryDciRegistrationRequest.SetClientToken(clientToken)

	retryDciRegistrationResponse, err := app.ModuleClients.BccrClient.RetryDciRegistration(retryDciRegistrationRequest)
	if err != nil {
		return nil, errCommon.ReturnError(err, msg.ErrRetryDciRegistration, "补正申请 失败: ")
	}

	errCommon.NoReturnInfo(retryDciRegistrationResponse, "补正申请 成功: ")

	_ = copier.CopyWithOption(&res, retryDciRegistrationResponse, copier.Option{DeepCopy: false})

	return
}

// CloseDciRegistration 数登停止申请
func (r *DciRegistration) CloseDciRegistration(req *dci.CloseDciRegistrationRequest) (res *dci.CloseDciRegistrationResponse, err error) {
	errCommon.NoReturnInfo(req, "数登停止申请 参数信息: ")

	res = new(dci.CloseDciRegistrationResponse)

	closeDciRegistrationRequest := new(bccrClient.CloseDciRegistrationRequest)
	_ = copier.CopyWithOption(&closeDciRegistrationRequest, req, copier.Option{DeepCopy: false})

	clientToken, err := createToken(time.Now().UnixMilli(), req.DigitalRegisterId, req.Name, req.MobileNo, app.ModuleClients.SfNode.Generate().Base64())
	if err != nil {
		return nil, errCommon.ReturnError(err, msg.ErrCreateClientToken, "创建clientToken 失败: ")
	}
	closeDciRegistrationRequest.SetClientToken(clientToken)

	closeDciRegistrationResponse, err := app.ModuleClients.BccrClient.CloseDciRegistration(closeDciRegistrationRequest)
	if err != nil {
		return nil, errCommon.ReturnError(err, msg.ErrCloseDciRegistration, "数登停止申请 失败: ")
	}

	errCommon.NoReturnInfo(closeDciRegistrationResponse, "数登停止申请 成功: ")

	_ = copier.CopyWithOption(&res, closeDciRegistrationResponse, copier.Option{DeepCopy: false})

	return
}

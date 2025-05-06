package logic

import (
	"chain-dci/pb/dci"
	"chain-dci/pkg/app"
	errCommon "chain-dci/pkg/err"
	"chain-dci/pkg/msg"
	bccrClient "github.com/antchain-openapi-sdk-go/bccr/client"
	"github.com/jinzhu/copier"
	"time"
)

type IDciWork interface {
	CreateDciPreregistration(req *dci.CreateDciPreregistrationRequest) (res *dci.CreateDciPreregistrationResponse, err error)
	QueryDciPreregistration(req *dci.QueryDciPreregistrationRequest) (res *dci.QueryDciPreregistrationResponse, err error)
	SubmitDciFeedback(req *dci.SubmitDciFeedbackRequest) (res *dci.SubmitDciFeedbackResponse, err error)
}

type DciWork struct {
}

// CreateDciPreregistration 作品 申领 DCI
func (w *DciWork) CreateDciPreregistration(req *dci.CreateDciPreregistrationRequest) (res *dci.CreateDciPreregistrationResponse, err error) {
	errCommon.NoReturnInfo(req, "作品 申领 DCI 参数信息: ")

	res = new(dci.CreateDciPreregistrationResponse)

	createDciPreregistrationRequest := new(bccrClient.CreateDciPreregistrationRequest)
	_ = copier.CopyWithOption(&createDciPreregistrationRequest, req, copier.Option{DeepCopy: false})

	clientToken, err := createToken(time.Now().UnixMilli(), req.DciUserId, req.WorkName, req.AuthorName, req.WorkFileId, app.ModuleClients.SfNode.Generate().Base64())
	if err != nil {
		return nil, errCommon.ReturnError(err, msg.ErrCreateClientToken, "创建clientToken 失败: ")
	}
	createDciPreregistrationRequest.SetClientToken(clientToken)

	createDciPreregistrationResponse, err := app.ModuleClients.BccrClient.CreateDciPreregistration(createDciPreregistrationRequest)
	if err != nil {
		return nil, errCommon.ReturnError(err, msg.ErrCreateDciPreregistration, "作品 申领 DCI 失败: ")
	}

	errCommon.NoReturnInfo(createDciPreregistrationResponse, "作品 申领 DCI 成功: ")

	_ = copier.CopyWithOption(&res, createDciPreregistrationResponse, copier.Option{DeepCopy: false})

	return
}

// QueryDciPreregistration 作品 查询 DCI
func (w *DciWork) QueryDciPreregistration(req *dci.QueryDciPreregistrationRequest) (res *dci.QueryDciPreregistrationResponse, err error) {
	errCommon.NoReturnInfo(req, "作品 查询 DCI 参数信息: ")

	res = new(dci.QueryDciPreregistrationResponse)

	queryDciPreregistrationRequest := new(bccrClient.QueryDciPreregistrationRequest)
	_ = copier.CopyWithOption(&queryDciPreregistrationRequest, req, copier.Option{DeepCopy: false})

	queryDciPreregistrationResponse, err := app.ModuleClients.BccrClient.QueryDciPreregistration(queryDciPreregistrationRequest)
	if err != nil {
		return nil, errCommon.ReturnError(err, msg.ErrQueryDciPreregistration, "作品 查询 DCI 失败: ")
	}

	errCommon.NoReturnInfo(queryDciPreregistrationResponse, "作品 查询 DCI 成功: ")

	_ = copier.CopyWithOption(&res, queryDciPreregistrationResponse, copier.Option{DeepCopy: false})

	return
}

// SubmitDciFeedback  作品 DCI  申诉
func (w *DciWork) SubmitDciFeedback(req *dci.SubmitDciFeedbackRequest) (res *dci.SubmitDciFeedbackResponse, err error) {
	errCommon.NoReturnInfo(req, "作品 DCI  申诉 参数信息: ")

	res = new(dci.SubmitDciFeedbackResponse)

	submitDciFeedbackRequest := new(bccrClient.SubmitDciFeedbackRequest)
	_ = copier.CopyWithOption(&submitDciFeedbackRequest, req, copier.Option{DeepCopy: false})

	clientToken, err := createToken(time.Now().UnixMilli(), req.ServiceId, req.ContactPhoneNumber, req.ContactName, req.FeedbackType, req.Message, app.ModuleClients.SfNode.Generate().Base64())
	if err != nil {
		return nil, errCommon.ReturnError(err, msg.ErrCreateClientToken, "创建clientToken 失败: ")
	}
	submitDciFeedbackRequest.SetClientToken(clientToken)

	submitDciFeedbackResponse, err := app.ModuleClients.BccrClient.SubmitDciFeedback(submitDciFeedbackRequest)
	if err != nil {
		return nil, errCommon.ReturnError(err, msg.ErrSubmitDciFeedback, "作品 DCI  申诉 失败: ")
	}

	errCommon.NoReturnInfo(submitDciFeedbackResponse, "作品 DCI  申诉 成功: ")

	_ = copier.CopyWithOption(&res, submitDciFeedbackResponse, copier.Option{DeepCopy: false})

	return
}

// QueryDciFeedback  作品 DCI  申诉 查询
func (w *DciWork) QueryDciFeedback(req *dci.QueryDciFeedbackRequest) (res *dci.QueryDciFeedbackResponse, err error) {
	errCommon.NoReturnInfo(req, "作品 DCI  申诉 参数信息: ")

	res = new(dci.QueryDciFeedbackResponse)

	submitDciFeedbackRequest := new(bccrClient.QueryDciFeedbackRequest)
	_ = copier.CopyWithOption(&submitDciFeedbackRequest, req, copier.Option{DeepCopy: false})

	queryDciFeedbackResponseResponse, err := app.ModuleClients.BccrClient.QueryDciFeedback(submitDciFeedbackRequest)
	if err != nil {
		return nil, errCommon.ReturnError(err, msg.ErrSubmitDciFeedback, "作品 DCI  申诉 查询 失败: ")
	}

	errCommon.NoReturnInfo(queryDciFeedbackResponseResponse, "作品 DCI  申诉 查询 成功: ")

	_ = copier.CopyWithOption(&res, queryDciFeedbackResponseResponse, copier.Option{DeepCopy: false})

	return
}

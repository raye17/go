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

type IDciUser interface {
	AddDciUser(req *dci.AddDciUserRequest) (res *dci.AddDciUserResponse, err error)
	UpdateDciUser(req *dci.UpdateDciUserRequest) (res *dci.UpdateDciUserResponse, err error)
	QueryDciUser(req *dci.QueryDciUserRequest) (res *dci.QueryDciUserResponse, err error)
}

type DciUser struct {
}

// AddDciUser 著作权人 申领 DCI信息
func (u *DciUser) AddDciUser(req *dci.AddDciUserRequest) (res *dci.AddDciUserResponse, err error) {
	errCommon.NoReturnInfo(req, "著作权人 申领 DCI信息 参数信息: ")

	res = new(dci.AddDciUserResponse)

	addDciUserRequest := new(bccrClient.AddDciUserRequest)
	_ = copier.CopyWithOption(&addDciUserRequest, req, copier.Option{DeepCopy: false})

	clientToken, err := createToken(time.Now().UnixMilli(), req.CertName, req.CertificateNumber, req.Phone, app.ModuleClients.SfNode.Generate().Base64())
	if err != nil {
		return nil, errCommon.ReturnError(err, msg.ErrCreateClientToken, "创建clientToken 失败: ")
	}
	addDciUserRequest.SetClientToken(clientToken)

	addDciUserResponse, err := app.ModuleClients.BccrClient.AddDciUser(addDciUserRequest)
	if err != nil {
		return nil, errCommon.ReturnError(err, msg.ErrAddDciUser, "著作权人 申领 DCI信息 失败: ")
	}

	errCommon.NoReturnInfo(addDciUserResponse, "著作权人 申领 DCI信息 成功: ")

	_ = copier.CopyWithOption(&res, addDciUserResponse, copier.Option{DeepCopy: false})

	return
}

// UpdateDciUser 著作权人 更新 DCI信息
func (u *DciUser) UpdateDciUser(req *dci.UpdateDciUserRequest) (res *dci.UpdateDciUserResponse, err error) {
	errCommon.NoReturnInfo(req, "著作权人 更新 DCI信息 参数信息: ")

	res = new(dci.UpdateDciUserResponse)

	updateDciUserRequest := new(bccrClient.UpdateDciUserRequest)
	_ = copier.CopyWithOption(&updateDciUserRequest, req, copier.Option{DeepCopy: false})

	clientToken, err := createToken(time.Now().UnixMilli(), req.DciUserId, req.CertFrontFileId, req.CertBackFileId, req.Phone, app.ModuleClients.SfNode.Generate().Base64())
	if err != nil {
		return nil, errCommon.ReturnError(err, msg.ErrCreateClientToken, "创建clientToken 失败: ")
	}
	updateDciUserRequest.SetClientToken(clientToken)

	updateDciUserResponse, err := app.ModuleClients.BccrClient.UpdateDciUser(updateDciUserRequest)
	if err != nil {
		return nil, errCommon.ReturnError(err, msg.ErrUpdateDciUser, "著作权人 更新 DCI信息 失败: ")
	}

	errCommon.NoReturnInfo(updateDciUserResponse, "著作权人 更新 DCI信息 成功: ")

	_ = copier.CopyWithOption(&res, updateDciUserResponse, copier.Option{DeepCopy: false})

	return
}

// QueryDciUser 著作权人 查询 DCI信息
func (u *DciUser) QueryDciUser(req *dci.QueryDciUserRequest) (res *dci.QueryDciUserResponse, err error) {
	errCommon.NoReturnInfo(req, "著作权人 查询 DCI信息 参数信息: ")

	res = new(dci.QueryDciUserResponse)

	queryDciUserRequest := new(bccrClient.QueryDciUserRequest)
	_ = copier.CopyWithOption(&queryDciUserRequest, req, copier.Option{DeepCopy: false})

	queryDciUserResponse, err := app.ModuleClients.BccrClient.QueryDciUser(queryDciUserRequest)
	if err != nil {
		return nil, errCommon.ReturnError(err, msg.ErrQueryDciUser, "著作权人 查询 DCI信息 失败: ")
	}

	errCommon.NoReturnInfo(queryDciUserResponse, "著作权人 查询 DCI信息 成功: ")

	_ = copier.CopyWithOption(&res, queryDciUserResponse, copier.Option{DeepCopy: false})

	return
}

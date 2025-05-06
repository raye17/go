package controller

import (
	"chain-dci/internal/logic"
	"chain-dci/pb/dci"
	errCommon "chain-dci/pkg/err"
	"chain-dci/pkg/msg"
	"context"
	"errors"
)

type DciProvider struct {
	dci.UnimplementedDciServer
	file    *logic.File
	dciUser *logic.DciUser
	dciWork *logic.DciWork
	dciReg  *logic.DciRegistration
}

func (d *DciProvider) GetUploadUrl(_ context.Context, req *dci.GetUploadUrlRequest) (res *dci.GetUploadUrlResponse, err error) {
	res = new(dci.GetUploadUrlResponse)
	if req.FileName == "" {
		return res, errCommon.ReturnError(errors.New(msg.ErrFileIsEmpty), msg.ErrFileIsEmpty, "未获取到相关文件 :")
	}

	res, err = d.file.Upload(req)
	return
}

func (d *DciProvider) AddDciUser(_ context.Context, req *dci.AddDciUserRequest) (res *dci.AddDciUserResponse, err error) {
	res = new(dci.AddDciUserResponse)
	res, err = d.dciUser.AddDciUser(req)
	return
}

func (d *DciProvider) UpdateDciUser(_ context.Context, req *dci.UpdateDciUserRequest) (res *dci.UpdateDciUserResponse, err error) {
	res = new(dci.UpdateDciUserResponse)
	res, err = d.dciUser.UpdateDciUser(req)
	return
}

func (d *DciProvider) QueryDciUser(_ context.Context, req *dci.QueryDciUserRequest) (res *dci.QueryDciUserResponse, err error) {
	res = new(dci.QueryDciUserResponse)
	res, err = d.dciUser.QueryDciUser(req)
	return
}

func (d *DciProvider) CreateDciPreregistration(_ context.Context, req *dci.CreateDciPreregistrationRequest) (res *dci.CreateDciPreregistrationResponse, err error) {
	res = new(dci.CreateDciPreregistrationResponse)
	res, err = d.dciWork.CreateDciPreregistration(req)
	return
}

func (d *DciProvider) QueryDciPreregistration(_ context.Context, req *dci.QueryDciPreregistrationRequest) (res *dci.QueryDciPreregistrationResponse, err error) {
	res = new(dci.QueryDciPreregistrationResponse)
	res, err = d.dciWork.QueryDciPreregistration(req)
	return
}

func (d *DciProvider) CreateDciRegistration(_ context.Context, req *dci.CreateDciRegistrationRequest) (res *dci.CreateDciRegistrationResponse, err error) {
	res = new(dci.CreateDciRegistrationResponse)
	res, err = d.dciReg.CreateDciRegistration(req)
	return
}

func (d *DciProvider) QueryDciRegistration(_ context.Context, req *dci.QueryDciRegistrationRequest) (res *dci.QueryDciRegistrationResponse, err error) {
	res = new(dci.QueryDciRegistrationResponse)
	res, err = d.dciReg.QueryDciRegistration(req)
	return
}

func (d *DciProvider) GetDciPayUrl(_ context.Context, req *dci.GetDciPayUrlRequest) (res *dci.GetDciPayUrlResponse, err error) {
	res = new(dci.GetDciPayUrlResponse)
	res, err = d.dciReg.GetDciPayUrl(req)
	return
}

func (d *DciProvider) GetDciPayAmount(_ context.Context, req *dci.GetDciPayUrlRequest) (res *dci.GetDciPayUrlResponse, err error) {
	res = new(dci.GetDciPayUrlResponse)
	res, err = d.dciReg.GetDciPayUrl(req)
	return
}

func (d *DciProvider) QueryDciPay(_ context.Context, req *dci.QueryDciPayRequest) (res *dci.QueryDciPayResponse, err error) {
	res = new(dci.QueryDciPayResponse)
	res, err = d.dciReg.QueryDciPay(req)
	return
}

func (d *DciProvider) GetDciRegistrationcert(_ context.Context, req *dci.GetDciRegistrationcertRequest) (res *dci.GetDciRegistrationcertResponse, err error) {
	res = new(dci.GetDciRegistrationcertResponse)
	res, err = d.dciReg.GetDciRegistrationcert(req)
	return
}

func (d *DciProvider) RetryDciRegistration(_ context.Context, req *dci.RetryDciRegistrationRequest) (res *dci.RetryDciRegistrationResponse, err error) {
	res = new(dci.RetryDciRegistrationResponse)
	res, err = d.dciReg.RetryDciRegistration(req)
	return
}

func (d *DciProvider) CloseDciRegistration(_ context.Context, req *dci.CloseDciRegistrationRequest) (res *dci.CloseDciRegistrationResponse, err error) {
	res = new(dci.CloseDciRegistrationResponse)
	res, err = d.dciReg.CloseDciRegistration(req)
	return
}

func (d *DciProvider) SubmitDciFeedback(_ context.Context, req *dci.SubmitDciFeedbackRequest) (res *dci.SubmitDciFeedbackResponse, err error) {
	res = new(dci.SubmitDciFeedbackResponse)
	res, err = d.dciWork.SubmitDciFeedback(req)
	return
}

func (d *DciProvider) QueryDciFeedback(_ context.Context, req *dci.QueryDciFeedbackRequest) (res *dci.QueryDciFeedbackResponse, err error) {
	res = new(dci.QueryDciFeedbackResponse)
	res, err = d.dciWork.QueryDciFeedback(req)
	return
}

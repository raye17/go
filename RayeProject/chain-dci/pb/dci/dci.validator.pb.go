// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pb/dci.proto

package dci

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/protobuf/types/descriptorpb"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *GetUploadUrlRequest) Validate() error {
	if this.FileName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("FileName", fmt.Errorf(`文件不能为空`))
	}
	return nil
}
func (this *GetUploadUrlResponse) Validate() error {
	return nil
}
func (this *AddDciUserRequest) Validate() error {
	if this.CertName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("CertName", fmt.Errorf(`证件名称不能为空`))
	}
	if this.CertificateNumber == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("CertificateNumber", fmt.Errorf(`证件号码不能为空`))
	}
	if this.CertificateType == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("CertificateType", fmt.Errorf(`证件类型不能为空`))
	}
	if this.CertificateFrontFileId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("CertificateFrontFileId", fmt.Errorf(`证件正面文件路径不能为空`))
	}
	if this.Phone == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Phone", fmt.Errorf(`手机号不能为空`))
	}
	if this.AreaType == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("AreaType", fmt.Errorf(`所属地区不能为空`))
	}
	return nil
}
func (this *AddDciUserResponse) Validate() error {
	return nil
}
func (this *UpdateDciUserRequest) Validate() error {
	if this.DciUserId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("DciUserId", fmt.Errorf(`dci用户ID不能为空`))
	}
	if this.CertFrontFileId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("CertFrontFileId", fmt.Errorf(`证件正面文件路径不能为空`))
	}
	return nil
}
func (this *UpdateDciUserResponse) Validate() error {
	return nil
}
func (this *QueryDciUserRequest) Validate() error {
	if this.CertificateType == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("CertificateType", fmt.Errorf(`证件类型不能为空`))
	}
	if this.CertificateNumber == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("CertificateNumber", fmt.Errorf(`证件号码不能为空`))
	}
	if this.Phone == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Phone", fmt.Errorf(`手机号不能为空`))
	}
	return nil
}
func (this *QueryDciUserResponse) Validate() error {
	return nil
}
func (this *CreateDciPreregistrationRequest) Validate() error {
	if this.WorkName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("WorkName", fmt.Errorf(`作品名称不能为空`))
	}
	if !(len(this.WorkName) < 51) {
		return github_com_mwitkow_go_proto_validators.FieldError("WorkName", fmt.Errorf(`作品名称不能为空`))
	}
	if this.DciUserId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("DciUserId", fmt.Errorf(`DCI用户ID不能为空`))
	}
	if this.WorkCategory == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("WorkCategory", fmt.Errorf(`作品类型不能为空`))
	}
	if this.WorkFileId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("WorkFileId", fmt.Errorf(`作品文件路径不能为空`))
	}
	if this.FileType == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("FileType", fmt.Errorf(`文件类型不能为空`))
	}
	if this.CreationInfo != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CreationInfo); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CreationInfo", err)
		}
	}
	if this.PublicationInfo != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.PublicationInfo); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("PublicationInfo", err)
		}
	}
	if this.AuthorName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("AuthorName", fmt.Errorf(`作者姓名不能为空`))
	}
	if this.AuthorSignature == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("AuthorSignature", fmt.Errorf(`作者署名不能为空`))
	}
	if this.RightInfo != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.RightInfo); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("RightInfo", err)
		}
	}
	if this.PreRegistrationTrueWill == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PreRegistrationTrueWill", fmt.Errorf(`真实意愿表达信息不能为空`))
	}
	for _, item := range this.CopyrightOwnerIds {
		if item == "" {
			return github_com_mwitkow_go_proto_validators.FieldError("CopyrightOwnerIds", fmt.Errorf(`著作权人用户id列表不能为空`))
		}
	}
	return nil
}
func (this *DciCreationInfo) Validate() error {
	if this.CreationNature == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("CreationNature", fmt.Errorf(`作品创作性质不能为空`))
	}
	if this.CreationCompletionDate == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("CreationCompletionDate", fmt.Errorf(`创作完成日期不能为空`))
	}
	if this.CreationCompletionCode == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("CreationCompletionCode", fmt.Errorf(`作品创作地点地区编码不能为空`))
	}
	return nil
}
func (this *DciPublicationInfo) Validate() error {
	if this.PublicationStatus == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PublicationStatus", fmt.Errorf(`作品发表状态不能为空`))
	}
	if this.FirstPublicationDate == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("FirstPublicationDate", fmt.Errorf(`首次发表日期不能为空`))
	}
	if this.FirstPublicationCode == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("FirstPublicationCode", fmt.Errorf(`首次发表地点地区编码不能为空`))
	}
	return nil
}
func (this *DciRightInfo) Validate() error {
	if this.RightScope == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("RightScope", fmt.Errorf(`作品权利范围不能为空`))
	}
	if this.RightObtainWay == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("RightObtainWay", fmt.Errorf(`权利取得方式不能为空`))
	}
	if this.OwnershipWay == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("OwnershipWay", fmt.Errorf(`权利归属方式不能为空`))
	}
	return nil
}
func (this *CreateDciPreregistrationResponse) Validate() error {
	return nil
}
func (this *QueryDciPreregistrationRequest) Validate() error {
	if this.DciContentId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("DciContentId", fmt.Errorf(`DCI内容ID不能为空`))
	}
	return nil
}
func (this *QueryDciPreregistrationResponse) Validate() error {
	return nil
}
func (this *CreateDciRegistrationRequest) Validate() error {
	if this.DciContentId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("DciContentId", fmt.Errorf(`DCI申领ID不能为空`))
	}
	if this.ExplanationInfo != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ExplanationInfo); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ExplanationInfo", err)
		}
	}
	if this.InvoiceInfo != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.InvoiceInfo); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("InvoiceInfo", err)
		}
	}
	if this.AdditionalFileInfo != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.AdditionalFileInfo); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("AdditionalFileInfo", err)
		}
	}
	return nil
}
func (this *DciExplanationInfo) Validate() error {
	if this.CreationPurpose == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("CreationPurpose", fmt.Errorf(`创作目的不能为空`))
	}
	if !(len(this.CreationPurpose) < 51) {
		return github_com_mwitkow_go_proto_validators.FieldError("CreationPurpose", fmt.Errorf(`创作目的不能为空`))
	}
	if this.CreationProcess == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("CreationProcess", fmt.Errorf(`创作过程不能为空`))
	}
	if !(len(this.CreationProcess) < 86) {
		return github_com_mwitkow_go_proto_validators.FieldError("CreationProcess", fmt.Errorf(`创作过程不能为空`))
	}
	if this.Originality == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Originality", fmt.Errorf(`独创性说明不能为空`))
	}
	if !(len(this.Originality) < 141) {
		return github_com_mwitkow_go_proto_validators.FieldError("Originality", fmt.Errorf(`独创性说明不能为空`))
	}
	if this.FontCopyright == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("FontCopyright", fmt.Errorf(`字体声明不能为空`))
	}
	return nil
}
func (this *InvoiceInfo) Validate() error {
	if this.InvoiceType == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("InvoiceType", fmt.Errorf(`发票类型不能为空`))
	}
	if this.InvoiceHeader == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("InvoiceHeader", fmt.Errorf(`发票抬头不能为空`))
	}
	return nil
}
func (this *AdditionalFileInfo) Validate() error {
	return nil
}
func (this *CreateDciRegistrationResponse) Validate() error {
	return nil
}
func (this *QueryDciRegistrationRequest) Validate() error {
	return nil
}
func (this *QueryDciRegistrationResponse) Validate() error {
	return nil
}
func (this *GetDciPayUrlRequest) Validate() error {
	if this.DigitalRegisterId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("DigitalRegisterId", fmt.Errorf(`数登申请ID不能为空`))
	}
	return nil
}
func (this *GetDciPayUrlResponse) Validate() error {
	return nil
}
func (this *QueryDciPayRequest) Validate() error {
	if this.DigitalRegisterId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("DigitalRegisterId", fmt.Errorf(`数登申请ID不能为空`))
	}
	return nil
}
func (this *QueryDciPayResponse) Validate() error {
	return nil
}
func (this *GetDciRegistrationcertRequest) Validate() error {
	if this.DigitalRegisterId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("DigitalRegisterId", fmt.Errorf(`数登申请ID不能为空`))
	}
	return nil
}
func (this *GetDciRegistrationcertResponse) Validate() error {
	return nil
}
func (this *RetryDciRegistrationRequest) Validate() error {
	if this.DigitalRegisterId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("DigitalRegisterId", fmt.Errorf(`数登申请ID不能为空`))
	}
	if this.ExplanationInfo != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ExplanationInfo); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ExplanationInfo", err)
		}
	}
	if this.AdditionalFileInfo != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.AdditionalFileInfo); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("AdditionalFileInfo", err)
		}
	}
	return nil
}
func (this *RetryDciRegistrationResponse) Validate() error {
	return nil
}
func (this *CloseDciRegistrationRequest) Validate() error {
	if this.DigitalRegisterId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("DigitalRegisterId", fmt.Errorf(`数登申请ID不能为空`))
	}
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`退费人名称不能为空`))
	}
	if !(len(this.Name) < 41) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`退费人名称不能为空`))
	}
	if this.MobileNo == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("MobileNo", fmt.Errorf(`联系手机号不能为空`))
	}
	return nil
}
func (this *CloseDciRegistrationResponse) Validate() error {
	return nil
}
func (this *SubmitDciFeedbackRequest) Validate() error {
	if this.ServiceId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ServiceId", fmt.Errorf(`业务ID`))
	}
	if this.ContactName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ContactName", fmt.Errorf(`联系人`))
	}
	if this.ContactPhoneNumber == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ContactPhoneNumber", fmt.Errorf(`联系电话`))
	}
	if this.Message == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Message", fmt.Errorf(`申诉原因`))
	}
	if this.FeedbackType == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("FeedbackType", fmt.Errorf(`反馈类型`))
	}
	return nil
}
func (this *SubmitDciFeedbackResponse) Validate() error {
	return nil
}
func (this *QueryDciFeedbackRequest) Validate() error {
	if this.Id == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`申诉ID`))
	}
	return nil
}
func (this *QueryDciFeedbackResponse) Validate() error {
	return nil
}

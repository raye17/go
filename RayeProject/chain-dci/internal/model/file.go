package model

import "gorm.io/gorm"

type FileInfo struct {
	gorm.Model
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId string `json:"req_msg_id,omitempty" gorm:"req_msg_id"`
	// 结果码，一般OK表示调用成功
	ResultCode string `json:"result_code,omitempty" gorm:"result_code"`
	// 异常信息的文本描述
	ResultMsg string `json:"result_msg,omitempty" gorm:"result_msg"`
	// 授权访问oss链接
	Url string `json:"url,omitempty" gorm:"column:url;type:varchar(4096)"`
	// OSS 文件id
	FileId string `json:"file_id,omitempty" gorm:"file_id"`
	// 文件链接
	FileUrl string `json:"file_url,omitempty" gorm:"column:file_url;type:varchar(4096)"`
	// 真实文件是否上传
	IsUpload int `json:"is_upload" gorm:"column:is_upload;type:int;default:1"`
}

func (f *FileInfo) GetTableName() string {
	return "file_info"
}

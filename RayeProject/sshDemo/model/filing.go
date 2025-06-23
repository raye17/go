package model

import (
	"encoding/json"

	"gorm.io/datatypes"
	"gorm.io/plugin/soft_delete"
)

type SecFilings struct {
	Id              int32                 `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	FilingKey       string                `gorm:"column:filing_key;type:varchar(128);not null;uniqueIndex:uk_filingkey_deleted,priority:1" json:"filingKey"`
	FilingDate      string                `json:"filingDate" gorm:"column:filing_date;type:varchar(32);comment:披露日期"`
	Form            string                `json:"form" gorm:"column:form;type:varchar(2048);comment:报告类型"`
	Description     string                `json:"description" gorm:"column:description;type:varchar(2048);comment:报告摘要"`
	FormDescription string                `json:"formDescription" gorm:"column:form_description;type:varchar(2048);comment:报告类型描述"`
	FileLink        string                `json:"fileLink" gorm:"column:file_link;type:varchar(2048);comment:文件链接"`
	DataFiles       datatypes.JSON        `json:"dataFiles" gorm:"column:data_files;type:text;comment:数据文件"`
	PdfFile         string                `json:"pdfFile" gorm:"column:pdf_file;type:varchar(2048);comment:PDF文件"`
	WordFile        string                `json:"wordFile" gorm:"column:word_file;type:varchar(2048);comment:WORD文件"`
	ExcelFile       string                `json:"excelFile" gorm:"column:excel_file;type:varchar(2048);comment:EXCEL文件"`
	Status          int32                 `json:"status" gorm:"column:status;type:int;default:1;comment:1:下架状态 2:上架"`
	Operator        string                `json:"operator" gorm:"column:operator;type:varchar(32);comment:操作人"`
	OperatorId      int32                 `json:"operatorId" gorm:"column:operator_id;type:int(11);comment:操作人ID"`
	CreatedAt       int64                 `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt       int64                 `gorm:"column:updated_at;autoCreateTime"`
	DeletedAt       soft_delete.DeletedAt `gorm:"uniqueIndex:uk_filingkey_deleted,priority:2"`
}
type RawFiling struct {
	FilingDate      string          `json:"filingDate"`
	Form            string          `json:"form"`
	Description     string          `json:"description"`
	FormDescription string          `json:"formDescription"`
	FileLink        string          `json:"fileLink"`
	DataFiles       json.RawMessage `json:"dataFiles"` // 保留原始 JSON
	Idx             int64           `json:"idx"`
}
type FormType struct {
	Id        int32  `gorm:"column:id;type:int(11);primaryKey;autoIncrement" json:"id"`
	FormType  string `json:"formType" gorm:"column:form_type;type:varchar(2048);unique;comment:报告类型"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt int64  `gorm:"column:updated_at;autoCreateTime"`
	DeletedAt soft_delete.DeletedAt
}

func (FormType) TableName() string {
	return "form_type" // 👈 和数据库中真实表名保持一致
}

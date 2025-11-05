package model

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/datatypes"
	"gorm.io/plugin/soft_delete"
)

type Leader struct {
	LeaderId   uint64 `gorm:"column:leader_id" json:"leaderId"`
	LeaderName string `gorm:"column:leader_name" json:"leaderName"`
}

// Report main:订单修改记录:report
type Report struct {
	ID        uint64                `gorm:"column:id;index:idx_deleted_at_id,priority:2" json:"ID"`
	Domain    *string               `gorm:"column:domain" json:"domain"`
	CreatedAt time.Time             `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time             `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;index:idx_deleted_at_id,priority:1" json:"deletedAt"`
	//WeekTime    time.Time             `gorm:"column:week_time" json:"weekTime"`
	EndDate      string          `gorm:"column:end_date" json:"endDate"`
	StartDate    string          `gorm:"column:start_date" json:"startDate"`
	EndDateDue   string          `gorm:"column:end_date_due" json:"endDateDue"`
	StartDateDue string          `gorm:"column:start_date_due" json:"startDateDue"`
	Type         int             `gorm:"column:type" json:"type"`
	Status       int             `gorm:"column:status" json:"status"`
	Site         string          `gorm:"column:site" json:"site"`
	SiteId       uint64          `gorm:"column:site_id" json:"siteId"`
	LeaderID     uint64          `gorm:"column:leader_id" json:"leaderID"`
	LeaderName   string          `gorm:"column:leader_name" json:"leaderName"`
	SalesVolume  decimal.Decimal `gorm:"column:sales_volume" json:"salesVolume"`
	Amount       decimal.Decimal `gorm:"column:amount;type:decimal(20,6)" json:"amount"`        //销售总额(团队本场业绩)
	DueAmount    decimal.Decimal `gorm:"column:due_amount;type:decimal(20,6)" json:"dueAmount"` //本场基数
	//ReportUsers    []*ReportUser   `gorm:"foreignKey:ReportID;" json:"reportUsers"`
	//ReportRead     []*ReportRead   `gorm:"foreignKey:ReportID;" json:"reportRead"`
	IncreaseAmount decimal.Decimal `gorm:"column:increase_amount;type:decimal(20,6)" json:"increaseAmount"` //增长金额
	IncreaseRate   decimal.Decimal `gorm:"column:increase_rate;type:decimal(20,6)" json:"increaseRate"`     //增长率
	MomRate        string          `gorm:"column:mom_rate" json:"momRate"`                                  //环比
	MomAmount      decimal.Decimal `gorm:"column:mom_amount;type:decimal(20,6)" json:"momAmount"`           //环比金额
	CanRevoke      bool            `gorm:"column:can_revoke;default:false" json:"canRevoke"`                // 是否可以撤销
	Color          string          `gorm:"column:color" json:"color"`                                       //业绩健康码
	Leaders        datatypes.JSON  `gorm:"column:leaders" json:"leaders"`                                   //总监
	Field          uint64          `gorm:"column:field" json:"field"`                                       //场期
}

func (m *Report) TableName() string {
	return "report"
}



package main

import (
	"fmt"
	"log"
	"strings"
	"study/db/common/db"
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/plugin/soft_delete"
)

type Record struct {
	ID        uint64                `gorm:"column:id" json:"ID"`
	Domain    *string               `gorm:"column:domain" json:"domain"`
	CreatedAt time.Time             `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time             `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"`
	Name      string                `gorm:"column:name" json:"name"`
	TelNum    string                `gorm:"column:tel_num" json:"telNum"`
	EndDate   string                `gorm:"column:end_date" json:"endDate"`
	StartDate string                `gorm:"column:start_date" json:"startDate"`
	Remark    string                `gorm:"column:remark" json:"remark"`
	Age       int                   `gorm:"column:age" json:"age"`
}

func main() {
	db, err := db.DbInit("test02")
	if err != nil {
		log.Println(err)
		return
	}
	db.AutoMigrate(&Record{})
	amount := decimal.NewFromInt(758000)
	preAmount := decimal.NewFromInt(606100)
	momAmount := amount.Sub(preAmount)
	r := calc(momAmount, preAmount).Mul(decimal.NewFromInt(100)).Round(1)
	fmt.Println(momAmount, r)
}
func decimalToPercent(d decimal.Decimal) string {
	if d.IsZero() {
		return "0"
	}

	// 先乘100
	rate := d.Mul(decimal.NewFromInt(100))

	// 保留2位小数并去掉多余0
	str := rate.StringFixed(2)
	str = strings.TrimRight(str, "0")
	str = strings.TrimRight(str, ".")

	return str + "%"
}
func calc(sumAMount, preAmount decimal.Decimal) decimal.Decimal {
	return sumAMount.Div(preAmount)
}

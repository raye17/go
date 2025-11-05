package service

import (
	"fmt"
	"test001/api/account"
	"test001/api/department"
	"test001/api/order"
	"test001/api/position"
	"test001/api/rule"

	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
)

var OrderProvider = new(order.OrderClientImpl)
var DepartmentProvider = new(department.DepartmentClientImpl)
var PositionProvider = new(position.PositionClientImpl)
var RuleProvider = new(rule.RuleClientImpl)
var AccountProvider = new(account.AccountClientImpl)

func init() {
	config.SetConsumerService(DepartmentProvider)
	config.SetConsumerService(PositionProvider)
	config.SetConsumerService(RuleProvider)
	config.SetConsumerService(AccountProvider)
	config.SetConsumerService(OrderProvider)
	if err := config.Load(); err != nil {
		panic(err)
	}
	fmt.Println("init success")

}
func Ser() {
	fmt.Println("service init")
}

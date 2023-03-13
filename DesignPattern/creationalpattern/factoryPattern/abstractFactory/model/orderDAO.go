package model

import "fmt"

// orderMainDAO 订单主记录

type OrderMainDAO interface {
	SaveOrderMain()
}

// OrderDetailDAO 订单详情记录
type OrderDetailDAO interface {
	SaveOrderDetail()
}
type DAOFactory interface {
	CreateOrderMainDAO() OrderMainDAO
	CreateOrderDetailDAO() OrderDetailDAO
}
type RDBMainDAO struct{}

func (*RDBMainDAO) SaveOrderMain() {
	fmt.Println("rdb main...")
}

type RDBDetailDAO struct{}

func (*RDBDetailDAO) SaveOrderDetail() {
	fmt.Println("rdb detail...")
}

type RDBDAOFactory struct {
}

func (*RDBDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &RDBMainDAO{}
}
func (*RDBDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &RDBDetailDAO{}
}

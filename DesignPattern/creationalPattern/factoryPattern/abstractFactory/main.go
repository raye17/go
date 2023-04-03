package main

import (
	"abstractFactory/model"
)

func main() {
	factory := model.RDBDAOFactory{}
	i := factory.CreateOrderMainDAO()
	i.SaveOrderMain()
}

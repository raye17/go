package main

import (
	_ "test001/pkg/common/filter"
	"test001/pkg/router"
	"test001/pkg/service"
)

func main() {
	//sshProd.PrSh()
	//sshProd.FormType()
	//locals.ImportSecFilings()
	//locals.FormType()
	//fiee.SoftDeleteVideo()
	//fiee.UpdateVideoTime()
	//fiee.InsertApproval()
	//fiee.JIan1()
	service.Ser()
	r := router.NewRouter()
	r.Run(":8099")
}

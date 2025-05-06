package main

import (
	dciConfig "chain-dci/config"
	"chain-dci/internal/controller"
	_ "chain-dci/internal/handler"
	"chain-dci/pkg/app"
	//common "chain-dci/pkg/init"
	"chain-dci/pkg/tracing"
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/filter/tps/strategy"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	bccrClient "github.com/antchain-openapi-sdk-go/bccr/client"
	"github.com/bwmarrin/snowflake"
	"go.uber.org/zap"
)

func NewApp(Lg *zap.Logger, JaegerTracer *tracing.JaegerProvider, BccrClient *bccrClient.Client, SfNode *snowflake.Node) *app.App {
	return &app.App{
		Lg:           Lg,
		JaegerTracer: JaegerTracer,
		BccrClient:   BccrClient,
		SfNode:       SfNode,
		//DciDB:        DciDB,
	}
}

func main() {
	var err error
	dciConfig.GetOptions()
	app.ModuleClients, err = InitApp()
	if err != nil {
		panic(err)
	}

	//l, err := net.Listen("tcp", ":8883")
	//if err != nil {
	//	fmt.Printf("failed to listen: %v", err)
	//	return
	//}

	//s := grpc.NewServer()                               // 创建gRPC服务器
	//dci.RegisterDciServer(s, &controller.DciProvider{}) // 在gRPC服务端注册服务
	// 启动服务
	//err = s.Serve(l)
	//注册服务
	config.SetProviderService(&controller.DciProvider{})
	if err = config.Load(); err != nil {
		panic(err)
	}
	select {}
}

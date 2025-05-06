package app

import (
	"chain-dci/pkg/tracing"
	bccrClient "github.com/antchain-openapi-sdk-go/bccr/client"
	"github.com/bwmarrin/snowflake"
	"go.uber.org/zap"
)

var ModuleClients *App

type App struct {
	Lg *zap.Logger
	//RedisClient  *redis.Client
	JaegerTracer *tracing.JaegerProvider
	//DciDB        *gorm.DB
	SfNode     *snowflake.Node
	BccrClient *bccrClient.Client
}

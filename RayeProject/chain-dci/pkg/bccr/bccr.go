package bccr

import (
	dciConfig "chain-dci/config"
	bccrClient "github.com/antchain-openapi-sdk-go/bccr/client"
	"github.com/google/wire"
)

var Provider = wire.NewSet(NewBccrClient)

func NewBccrClient() *bccrClient.Client {
	// Endpoint 请参考 https://api.aliyun.com/product/rtc

	bccrConfig := new(bccrClient.Config)
	bccrConfig.SetEndpoint(dciConfig.Data.BccrYXT.EndPoint)
	bccrConfig.SetAccessKeyId(dciConfig.Data.BccrYXT.AccessKeyID)
	bccrConfig.SetAccessKeySecret(dciConfig.Data.BccrYXT.AccessKeySecret)

	client, err := bccrClient.NewClient(bccrConfig)
	if err != nil {
		panic(err)
	}
	return client
}

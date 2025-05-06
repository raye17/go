// go:build wireinject
//go:build wireinject
// +build wireinject

package main

import (
	"chain-dci/pkg/app"
	"chain-dci/pkg/bccr"
	"chain-dci/pkg/logger"
	"chain-dci/pkg/snowf"
	"chain-dci/pkg/tracing"
	"github.com/google/wire"
)

func InitApp() (*app.App, error) {
	wire.Build(logger.Provider, tracing.Provider, bccr.Provider, snowf.Provider, NewApp)
	return &app.App{}, nil
}

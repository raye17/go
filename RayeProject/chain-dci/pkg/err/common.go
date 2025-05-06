package err

import (
	"chain-dci/pkg/app"
	"errors"
	"fmt"
	"go.uber.org/zap"
)

func ReturnError(err error, msg, print string) error {
	if err != nil {
		field := zap.Field{}
		field.String = err.Error()
		app.ModuleClients.Lg.Error(print, field)
		fmt.Printf(print+"%+v\n", err)
		return errors.New(msg)
	}
	return nil
}

func NoReturnError(err error, print string) {
	if err != nil {
		field := zap.Field{}
		field.String = err.Error()
		app.ModuleClients.Lg.Error(print, field)
		fmt.Printf(print+"%+v\n", err)
	}
}

func NoReturnInfo(info interface{}, print string) {
	field := zap.Any("", info)
	app.ModuleClients.Lg.Info(print, field)
}

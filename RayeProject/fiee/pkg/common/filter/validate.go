package filter

import (
	"context"
	"fmt"
	"regexp"
	"strconv"

	"dubbo.apache.org/dubbo-go/v3/common/extension"
	"dubbo.apache.org/dubbo-go/v3/filter"
	"dubbo.apache.org/dubbo-go/v3/protocol"
	perrors "github.com/pkg/errors"
)

type validator interface {
	Validate() error
}

func init() {
	extension.SetFilter("fonValidateFilter", NewClientFonValidateFilter)
}

func NewClientFonValidateFilter() filter.Filter {
	return &ClientFonValidateFilter{}
}

type ClientFonValidateFilter struct {
}

func (f *ClientFonValidateFilter) Invoke(ctx context.Context, invoker protocol.Invoker, invocation protocol.Invocation) protocol.Result {

	if len(invocation.Arguments()) > 0 {
		if v, ok := invocation.Arguments()[0].(validator); ok {
			if err := v.Validate(); err != nil {
				errMsg := err.Error()
				re3, _ := regexp.Compile(`^invalid(.*): `)
				rep := re3.ReplaceAllString(errMsg, "")
				if errCode, err := strconv.ParseInt(rep, 10, 64); err == nil {
					return &protocol.RPCResult{Err: perrors.Errorf("%v", fmt.Sprintf("%d: %v", errCode, rep))}
				}

				return &protocol.RPCResult{Err: perrors.Errorf("%v", rep)}
			}
		}
	}

	return invoker.Invoke(ctx, invocation)
}

func (f *ClientFonValidateFilter) OnResponse(ctx context.Context, result protocol.Result, invoker protocol.Invoker, protocol protocol.Invocation) protocol.Result {
	return result
}

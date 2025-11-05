package filter

import (
	"context"
	"reflect"

	"dubbo.apache.org/dubbo-go/v3/common/extension"
	"dubbo.apache.org/dubbo-go/v3/filter"
	"dubbo.apache.org/dubbo-go/v3/protocol"
)

func init() {
	extension.SetFilter("fonDomainFilter", NewDomainFonFilter)
}

func NewDomainFonFilter() filter.Filter {
	return &DomainFonFilter{}
}

type DomainFonFilter struct {
}

func (f *DomainFonFilter) Invoke(ctx context.Context, invoker protocol.Invoker, invocation protocol.Invocation) protocol.Result {
	if len(invocation.Arguments()) > 0 {
		req := invocation.Arguments()[0]
		pp := reflect.ValueOf(req)
		field := pp.Elem().FieldByName("Domain")
		if field.IsValid() && field.String() == "" {
			field.SetString("fontree")
		}
		//if field.IsValid() {
		//	field.SetString(appconfig.AppConfig.Service.Domain)
		//}
	}

	return invoker.Invoke(ctx, invocation)
}

func (f *DomainFonFilter) OnResponse(ctx context.Context, result protocol.Result, invoker protocol.Invoker, protocol protocol.Invocation) protocol.Result {
	return result
}

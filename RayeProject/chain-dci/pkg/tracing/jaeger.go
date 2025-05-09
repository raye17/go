package tracing

import (
	dciConfig "chain-dci/config"
	"io"
	"time"

	"github.com/google/wire"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegerConfig "github.com/uber/jaeger-client-go/config"
	"go.uber.org/zap"
)

var Provider = wire.NewSet(NewTracing)

type JaegerProvider struct {
	Tracer opentracing.Tracer
	Closer io.Closer
}

//var JaegerPoint *JaegerProvider

func NewTracing() (jaegerProvider *JaegerProvider) {
	if dciConfig.Data.Jaeger.Open != "true" {
		return
	}
	jaegerProvider = &JaegerProvider{}
	cfg := jaegerConfig.Configuration{
		ServiceName: "oa-meeting",
		Sampler: &jaegerConfig.SamplerConfig{
			Type:  jaeger.SamplerTypeRemote,
			Param: 1,
		},
		Reporter: &jaegerConfig.ReporterConfig{
			LocalAgentHostPort:  dciConfig.Data.Jaeger.Addr,
			LogSpans:            true,
			BufferFlushInterval: 5 * time.Second,
		},
	}
	nativeTracerIo, closerIo, err := cfg.NewTracer(jaegerConfig.Logger(jaeger.StdLogger))
	if err != nil {
		zap.L().Error("nativeTracer err", zap.Error(err))
		return
	}
	opentracing.SetGlobalTracer(nativeTracerIo)
	jaegerProvider.Tracer = nativeTracerIo
	jaegerProvider.Closer = closerIo
	//JaegerPoint = jaegerProvider
	return
}

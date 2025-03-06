package gorm

import (
	"context"
	"log"
	"time"

	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"

	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"

	"go.opentelemetry.io/otel"
)

const (
	serviceName    = "GORM-Jaeger-Demo"
	jaegerEndpoint = "127.0.0.1:4318"
)

var tracer = otel.Tracer("gorm-demo")

func Trace01() {
	ctx := context.Background()
	tp, err := initTracer(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := tp.Shutdown(ctx); err != nil {
			log.Printf("error shutting down TracerProvider: %v", err)
		}
	}()
	dsn := "root:123456@tcp(127.0.0.1:3306)/test01?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Use(tracing.NewPlugin(tracing.WithoutMetrics())); err != nil {
		panic(err)
	}
	ctx, span := tracer.Start(ctx, "doSomething")
	defer span.End()
	if err := doSomething(ctx, db); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}
}

func initTracer(ctx context.Context) (*sdktrace.TracerProvider, error) {
	tp, err := newJaegerTraceProvider(ctx)
	if err != nil {
		return nil, err
	}
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp, nil
}

func newJaegerTraceProvider(ctx context.Context) (*sdktrace.TracerProvider, error) {
	exp, err := otlptracehttp.New(ctx,
		otlptracehttp.WithEndpoint(jaegerEndpoint),
		otlptracehttp.WithInsecure())
	if err != nil {
		return nil, err
	}
	res, err := resource.New(ctx, resource.WithAttributes(semconv.ServiceName(serviceName)))
	if err != nil {
		return nil, err
	}
	traceProvider := sdktrace.NewTracerProvider(
		sdktrace.WithResource(res),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exp, sdktrace.WithBatchTimeout(time.Second)),
	)
	return traceProvider, nil
}

func doSomething(ctx context.Context, db *gorm.DB) error {
	type Book struct {
		gorm.Model
		Title  string
		Author string
	}
	db.AutoMigrate(&Book{})
	db.WithContext(ctx).Create(&Book{Title: "go学习之路1", Author: "sxy"})
	var book Book
	if err := db.WithContext(ctx).Take(&book).Error; err != nil {
		return err
	}
	if err := db.WithContext(ctx).Delete(&book).Error; err != nil {
		return err
	}
	log.Println("done!")
	return nil
}

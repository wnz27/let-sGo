/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2023-03-09 14:21:33
 * @LastEditTime: 2023-03-09 15:06:55
 * @FilePath: /let-sGo/c/trace_learn/demos/baggage/other/otel_tttt.go
 * @description: type some description
 */
package other

import (
	"context"
	"log"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func InitTraceProvider(name string) func() {
	ctx := context.Background()

	// env := "local"
	// sn := "【" + env + "】" + "-" + name

	// otelAgentAddr := "127.0.0.1:4138"
	traceClientHttp := otlptracehttp.NewClient(
		// otlptracehttp.WithEndpoint(otelAgentAddr), // Endpoint 需替换为前提条件中获取的接入点信息。
		//otlptracehttp.WithURLPath("/adapt_xxxxx/api/otlp/traces"), //URLPath需替换为前提条件中获取的接入点信息。
		otlptracehttp.WithInsecure())
	otlptracehttp.WithCompression(1)

	traceExp, err := otlptrace.New(ctx, traceClientHttp)
	if err != nil {
		log.Println("Failed to create the collector trace exporter", err)
	}

	res, err := resource.New(ctx,
		resource.WithFromEnv(),
		resource.WithProcess(),
		resource.WithTelemetrySDK(),
		resource.WithHost(),
		resource.WithAttributes(
			// 在链路追踪后端显示的服务名称。
			semconv.ServiceNameKey.String(name),
		),
	)
	if err != nil {
		log.Println("failed to create resource", err)
	}

	bsp := sdktrace.NewBatchSpanProcessor(traceExp)
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(bsp),
	)

	// 设置全局propagator为tracecontext（默认不设置）。
	otel.SetTextMapPropagator(propagation.TraceContext{})
	otel.SetTracerProvider(tracerProvider)

	log.Println("OTEL init success", "init OTEL", "OK")

	return func() {
		cxt, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()
		if err := traceExp.Shutdown(cxt); err != nil {
			otel.Handle(err)
		}
	}
}

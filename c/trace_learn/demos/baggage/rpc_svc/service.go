/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2023-03-09 12:07:49
 * @LastEdiTime: 2023-03-09 14:30:48
 * @FilePath: /let-sGo/c/trace_learn/demos/baggage/rpc_svc/service.go
 * @description: type some description
 */

package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	h_v1 "fzkprac/c/trace_learn/demos/proto_gen/t1/v1"

	// "codeup.aliyun.com/63a12bb98d9a873a30aad6aa/LBM/mock-demo/observability/tracing"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func extractTraceId(ctx context.Context, fullId string, logger log.Logger) context.Context {
	ids := strings.Split(fullId, ":")
	traceId, spanId := ids[0], ids[1]
	if len(traceId) < 32 {
		traceId = "0000000000000000" + traceId
	}
	tid, err := trace.TraceIDFromHex(traceId)
	if err != nil {
		log.Println(err)
		return ctx
	}
	sid, err := trace.SpanIDFromHex(spanId)
	if err != nil {
		log.Println(err)
		return ctx
	}
	spanContext := trace.NewSpanContext(trace.SpanContextConfig{
		SpanID:     sid,
		TraceID:    tid,
		TraceFlags: 1,
		TraceState: trace.TraceState{},
		Remote:     true,
	})
	ctx = trace.ContextWithSpanContext(ctx, spanContext)
	return ctx
}

type serverInfo struct {
	Type string
	Info interface{}
}

const traceIdKey = "uber-trace-id"

func ContextToGRPC(tp trace.TracerProvider, logger log.Logger) func(ctx context.Context, md *metadata.MD) context.Context {
	return func(ctx context.Context, md *metadata.MD) context.Context {
		if span := trace.SpanFromContext(ctx); span != nil {
			spanContext := span.SpanContext()
			traceId := spanContext.TraceID
			spanId := spanContext.SpanID
			md.Set(traceIdKey, fmt.Sprintf("%s:%s:%s:%s", traceId().String(), spanId().String(), "00000000", "1"))
		}
		return ctx
	}
}

func GRPCToContext(tp trace.TracerProvider, operationName string, logger log.Logger) func(ctx context.Context, md metadata.MD) context.Context {
	return func(ctx context.Context, md metadata.MD) context.Context {
		metadata := md.Get(traceIdKey)
		ctx = context.WithValue(ctx, serverInfo{}, serverInfo{Type: "GRPC"})
		if metadata == nil {
			return ctx
		}
		return extractTraceId(ctx, metadata[0], logger)
	}
}

type H1Service struct {
	h_v1.UnimplementedHelloServiceServer
}

func NewH1Service() *H1Service {
	return &H1Service{}
}

func (srv *H1Service) Hello111(ctx context.Context, req *h_v1.Req1) (*h_v1.Res1, error) {
	md1, exist1 := metadata.FromIncomingContext(ctx)
	fmt.Println("-----<>>", exist1)
	tidmd := md1.Get("tid")
	fmt.Println(" ==================================== ", tidmd)

	tracer := otel.GetTracerProvider().Tracer("local-rpc")
	ctx, span := tracer.Start(ctx, "Hello111")
	defer span.End()
	// aa := propagation.Baggage{}
	// c1 := propagation.MapCarrier{}
	// aa.Inject(ctx, c1)

	// reqCtx := baggage.ContextWithBaggage(ctx, bag)

	return &h_v1.Res1{
		// B1: "【" + c1.Get("tid") + "】" + "---- done",
		B1: "【" + tidmd[0] + "】" + "---- done",
	}, nil
}

func (s *H1Service) String() string {
	return "lalala-local-demo"
}

func (s *H1Service) Register(srv *grpc.Server) {
	h_v1.RegisterHelloServiceServer(srv, s)
}

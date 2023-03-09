/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2023-03-09 11:51:43
 * @LastEditTime: 2023-03-09 15:50:46
 * @FilePath: /let-sGo/c/trace_learn/demos/baggage/http_svc/main.go
 * @description: type some description
 */

package main

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	// opt "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin" // http
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/baggage"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"fzkprac/c/trace_learn/demos/baggage/other"
	h_v1 "fzkprac/c/trace_learn/demos/proto_gen/t1/v1"
)

func lalala(c *gin.Context) {
	tracerP := otel.GetTracerProvider()
	tracer := tracerP.Tracer("local-demo")

	conn, dialErr := grpc.Dial("localhost:8099",
		grpc.WithInsecure(),
		// grpc.WithUnaryInterceptor(opt.UnaryClientInterceptor()),
	)
	if dialErr != nil {
		log.Println("rpcclt count pushing did not connect", dialErr)
		return
	}
	defer func() {
		if connErr := conn.Close(); connErr != nil {
			log.Println("conn err", connErr)
		}
	}()
	conn1 := h_v1.NewHelloServiceClient(conn)

	tidParam, _ := baggage.NewMember("tid", "tid_12345678")
	bag, _ := baggage.New(tidParam)
	reqCtx := baggage.ContextWithBaggage(c.Request.Context(), bag)
	ctx, span := tracer.Start(reqCtx, c.FullPath())
	defer span.End()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	r, receiveCountErr := conn1.Hello111(ctx, &h_v1.Req1{
		A1: "xxx",
	})
	if receiveCountErr != nil {
		log.Println("rpcclt h111", zap.Error(receiveCountErr))
		return
	}

	log.Println("----->", r.B1)
}

func main() {
	other.InitTraceProvider("local-http")
	// 新建一个没有任何默认中间件的路由
	r := gin.New()

	// 全局中间件
	// Logger 中间件将日志写入 gin.DefaultWriter，即使你将 GIN_MODE 设置为 release。
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery 中间件会 recover 任何 panic。如果有 panic 的话，会写入 500。
	r.Use(gin.Recovery())

	// 你可以为每个路由添加任意数量的中间件。
	r.GET("/hello-test", otelgin.Middleware("http-local-demo"), lalala)

	// 认证路由组
	// authorized := r.Group("/", AuthRequired())
	// 和使用以下两行代码的效果完全一样:
	// authorized := r.Group("/")
	// 路由组中间件! 在此例中，我们在 "authorized" 路由组中使用自定义创建的
	// AuthRequired() 中间件
	// authorized.Use(AuthRequired())
	// {
	// 	authorized.POST("/login", loginEndpoint)
	// 	authorized.POST("/submit", submitEndpoint)
	// 	authorized.POST("/read", readEndpoint)

	// 	// 嵌套路由组
	// 	testing := authorized.Group("testing")
	// 	testing.GET("/analytics", analyticsEndpoint)
	// }

	// 监听并在 0.0.0.0:80888 上启动服务
	r.Run(":8088")
}

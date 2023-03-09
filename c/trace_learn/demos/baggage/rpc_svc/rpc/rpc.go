/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2023-03-09 12:34:58
 * @LastEditTime: 2023-03-09 16:30:22
 * @FilePath: /let-sGo/c/trace_learn/demos/baggage/rpc_svc/rpc/rpc.go
 * @description: type some description
 */

package rpc

import (
	"fmt"
	"net"

	// "go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin" // http
	// "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	// "go.opentelemetry.io/otel"

	// "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	// opt "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type ServerInfo struct {
	Name string `json:"name"`
	Host string `json:"host"`
	Port int    `json:"port"`
}

type Server struct {
	Port     int
	rpc      *grpc.Server
	listener net.Listener
}

func NewRPCServer(host string, port int) *Server {
	srv := &Server{
		Port: port,
	}
	var err error
	domain := fmt.Sprintf("%s:%d", host, port)
	srv.listener, err = net.Listen("tcp", domain)
	if err != nil {
		panic(err)
	}

	rpcOpts := []grpc.ServerOption{
		// otelgin.Middleware("local-demo-trace"),
		// otelgrpc.WithTracerProvider(otel.GetTracerProvider()),
		// grpc.UnaryInterceptor(opt.UnaryServerInterceptor()),
	}
	srv.rpc = grpc.NewServer(rpcOpts...)

	return srv
}

func (srv *Server) Register() {
	reflection.Register(srv.rpc)
}

func (srv *Server) RServer() *grpc.Server {
	return srv.rpc
}

func (srv *Server) Run() {
	if err := srv.rpc.Serve(srv.listener); err != nil {
		panic(err)
	}
}

func (srv *Server) Stop() {
	srv.rpc.Stop()
}

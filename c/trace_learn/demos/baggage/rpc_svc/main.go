/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2023-03-09 11:53:04
 * @LastEditTime: 2023-03-09 15:07:23
 * @FilePath: /let-sGo/c/trace_learn/demos/baggage/rpc_svc/main.go
 * @description: type some description
 */

package main

import (
	"fzkprac/c/trace_learn/demos/baggage/other"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	rand.Seed(time.Now().UTC().UnixNano())
	myApp := New()
	RegisterRPC()

	// stdTracer := jaegerImpl.NewTracer(
	// 	myApp.Name(), myApp.Env(),
	// 	myApp.TraceConfig(),
	other.InitTraceProvider("local-rpc")

	myApp.RPCServer().Run()
	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop
	myApp.Stop()
}

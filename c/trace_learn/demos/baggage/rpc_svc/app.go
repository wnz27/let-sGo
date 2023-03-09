/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2023-03-09 12:36:39
 * @LastEditTime: 2023-03-09 12:54:57
 * @FilePath: /let-sGo/c/trace_learn/demos/baggage/rpc_svc/app.go
 * @description: type some description
 */
package main

import (
	selfRpc "fzkprac/c/trace_learn/demos/baggage/rpc_svc/rpc"
	"time"
)

type App struct {
	rpcServers []*selfRpc.Server
}

func New() *App {
	app := &App{
		rpcServers: make([]*selfRpc.Server, 0),
	}
	return app
}

// rpc服务
func (app *App) RPCServer() *App {
	app.createRPCServer()
	return app
}

func (app *App) createRPCServer() {
	baseHost := "localhost"
	basePort := 8099

	srv := selfRpc.NewRPCServer(baseHost, basePort)
	all := selfRpc.All()
	for _, sv := range all {
		sv.Register(srv.RServer())
	}
	srv.Register()
	app.rpcServers = append(app.rpcServers, srv)

}

func (app *App) Run() {
	// go func() {
	// 	pprfErr := http.ListenAndServe(":6060", nil)
	// 	runtimeApp.RuntimeLogger.Error("pprof-error", zap.Error(pprfErr))
	// }()
	app.run()
}

func (app *App) Stop() {
	app.stop()
}

func (app *App) stop() {
	// app.beforeStop
	// if app.rpcServer != nil {
	// 	app.rpcServer.Stop()
	// }
	for _, srv := range app.rpcServers {
		srv.Stop()
	}
	//app.afterAppStop()
}

func (app *App) run() {
	// app.beforeRun
	// if app.rpcServer != nil {
	// 	go app.rpcServer.Run()
	// }
	// todo 这里应该可以优化下 groutine 管理
	_, _ = time.LoadLocation("Asia/Shanghai")
	for _, srv := range app.rpcServers {
		go srv.Run()
	}
	// app.afterRun
}

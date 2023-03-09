/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2023-03-09 12:45:46
 * @LastEditTime: 2023-03-09 13:07:44
 * @FilePath: /let-sGo/c/trace_learn/demos/baggage/rpc_svc/route.go
 * @description: type some description
 */

package main

import "fzkprac/c/trace_learn/demos/baggage/rpc_svc/rpc"

// Server 控制机服务
var Server = &rpc.RPC{
	Services: []rpc.Service{
		&H1Service{},
	},
}

func RegisterRPC() {
	rpc.Register("local-demo", Server)
}

/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2023-03-09 12:35:27
 * @LastEditTime: 2023-03-09 12:35:29
 * @FilePath: /let-sGo/c/trace_learn/demos/baggage/rpc_svc/rpc/rpc_services.go
 * @description: type some description
 */

package rpc

import (
	"errors"
	"fmt"

	"google.golang.org/grpc"
)

type Service interface {
	String() string
	Register(srv *grpc.Server)
}
type RPC struct {
	Services []Service
}

func (rpc *RPC) Register(srv *grpc.Server) {
	for _, _s := range rpc.Services {
		_s.Register(srv)
	}
}

var servers = make(map[string]*RPC)

func New(name string) *RPC {
	return new(name)
}
func new(name string) *RPC {
	srv, ok := servers[name]
	if !ok {
		panic(fmt.Errorf("Unknown Server: %s", name))
	}

	return srv
}

func Register(name string, srv *RPC) {
	register(name, srv)
}
func register(name string, srv *RPC) {
	if _, ok := servers[name]; ok {
		panic(fmt.Sprintf("Server %s had registered", name))
	}

	servers[name] = srv
}

func GetRPC(name string) (*RPC, error) {
	srv, ok := servers[name]
	if !ok {
		return srv, errors.New(fmt.Sprintf("Unknow Server: %s", name))
	}

	return srv, nil
}

func All() []*RPC {
	all := make([]*RPC, 0)
	for _, srv := range servers {
		all = append(all, srv)
	}
	return all
}

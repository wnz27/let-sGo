/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/7/20 4:24 下午
* Description:
 */
package main

import (
	"context"
	"net/http"
	"time"
)

func handle(w http.ResponseWriter, req *http.Request) {
	// parent context
	timeout, _ := time.ParseDuration(req.FormValue("timeout"))
	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	// chidren context
	newCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	// do something...
}
/*
一般会有父级 context 和子级 context 的区别，我们要保证在程序的行为中上下文对于多个 goroutine 同时使用是安全的。
并且存在父子级别关系，父级 context 关闭或超时，可以继而影响到子级 context 的程序。
 */

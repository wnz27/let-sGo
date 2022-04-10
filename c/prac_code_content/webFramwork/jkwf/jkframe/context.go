/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2022-01-24 17:42:21
 * @LastEditTime: 2022-01-26 16:30:49
 * @FilePath: /let-sGo/prac_code_content/webFramwork/jkwf/jkframe/context.go
 * @description: type some description
 */

package jkframe

import (
	"context"
	"net/http"
	"strconv"
	"sync"
	"time"
)

/*
自己封装的 Context 最终需要提供四类功能函数：
- base 封装基本的函数功能，比如获取 http.Request 结构
- context 实现标准 Context 接口
- request 封装了 http.Request 的对外接口
- response 封装了 http.ResponseWriter 对外接口


在 Golang 中，每个 Goroutine 创建的时候，
我们要使用 defer 和 recover 关键字为当前 Goroutine 捕获 panic 异常，并进行处理，
否则，任意一处 panic 就会导致整个进程崩溃！
*/

type Context struct {
	// 获取请求、返回结果功能
	request        *http.Request
	responseWriter http.ResponseWriter
	ctx            context.Context
	// 当前请求的handler链条
	handlers []ControllerHandler
	index    int // 当前请求调用到调用链的哪个节点

	// 写保护机制
	writerMux *sync.Mutex
	// 是否超时标记
	hasTimeout bool

	params map[string]string // url路由匹配的参数

}

// 为context设置handlers
func (ctx *Context) SetHandlers(handlers []ControllerHandler) {
	ctx.handlers = handlers
}

// 核心函数，调用context的下一个函数
func (ctx *Context) Next() error {
	ctx.index++
	if ctx.index < len(ctx.handlers) {
		if err := ctx.handlers[ctx.index](ctx); err != nil {
			return err
		}
	}
	return nil
}

// NewContext 初始化一个Context
func NewContext(r *http.Request, w http.ResponseWriter) *Context {
	return &Context{
		request:        r,
		responseWriter: w,
		ctx:            r.Context(),
		writerMux:      &sync.Mutex{},
		index:          -1,
	}
}

// 设置参数
func (ctx *Context) SetParams(params map[string]string) {
	ctx.params = params
}

func (ctx *Context) GetRequest() *http.Request {
	return ctx.request
}

func (ctx *Context) GetResponse() http.ResponseWriter {
	return ctx.responseWriter
}

func (ctx *Context) SetHasTimeout() {
	ctx.hasTimeout = true
}

func (ctx *Context) HasTimeout() bool {
	return ctx.hasTimeout
}

// #region implement context.Context
func (ctx *Context) Deadline() (deadline time.Time, ok bool) {
	return ctx.BaseContext().Deadline()
}

// 对外暴露锁
func (ctx *Context) WriterMux() *sync.Mutex {
	return ctx.writerMux
}

func (ctx *Context) BaseContext() context.Context {
	return ctx.request.Context()
}

func (ctx *Context) Err() error {
	return ctx.BaseContext().Err()
}

func (ctx *Context) Done() <-chan struct{} {
	return ctx.BaseContext().Done()
}

func (ctx *Context) Value(key interface{}) interface{} {
	return ctx.BaseContext().Value(key)
}

func (ctx *Context) QueryArray(key string, def []string) []string {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		return vals
	}
	return def
}

// #endregion

// #region form post
func (ctx *Context) FormInt(key string, def int) int {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		len := len(vals)
		if len > 0 {
			intval, err := strconv.Atoi(vals[len-1])
			if err != nil {
				return def
			}
			return intval
		}
	}
	return def
}

func (ctx *Context) FormString(key string, def string) string {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		len := len(vals)
		if len > 0 {
			return vals[len-1]
		}
	}
	return def
}

func (ctx *Context) FormArray(key string, def []string) []string {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		return vals
	}
	return def
}

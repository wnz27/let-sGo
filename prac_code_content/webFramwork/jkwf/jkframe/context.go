/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2022-01-24 17:42:21
 * @LastEditTime: 2022-01-26 15:36:31
 * @FilePath: /let-sGo/prac_code_content/webFramwork/jkwf/jkframe/context.go
 * @description: type some description
 */

package jkframe

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
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

func NewContext(r *http.Request, w http.ResponseWriter) *Context {
	return &Context{
		request:        r,
		responseWriter: w,
		ctx:            r.Context(),
		writerMux:      &sync.Mutex{},
		index:          -1,
	}
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

// #endregion

// #region query url
func (ctx *Context) QueryInt(key string, def int) int {
	params := ctx.QueryAll()
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

func (ctx *Context) QueryString(key string, def string) string {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		len := len(vals)
		if len > 0 {
			return vals[len-1]
		}
	}
	return def
}

func (ctx *Context) QueryArray(key string, def []string) []string {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		return vals
	}
	return def
}

func (ctx *Context) QueryAll() map[string][]string {
	if ctx.request != nil {
		return map[string][]string(ctx.request.URL.Query())
	}
	return map[string][]string{}
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

func (ctx *Context) FormAll() map[string][]string {
	if ctx.request != nil {
		return map[string][]string(ctx.request.PostForm)
	}
	return map[string][]string{}
}

// #endregion

// #region application/json post

func (ctx *Context) BindJson(obj interface{}) error {
	if ctx.request != nil {
		body, err := ioutil.ReadAll(ctx.request.Body)
		if err != nil {
			return err
		}
		ctx.request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		err = json.Unmarshal(body, obj)
		if err != nil {
			return err
		}
	} else {
		return errors.New("ctx.request empty")
	}
	return nil
}

// #endregion

// #region response

func (ctx *Context) Json(status int, obj interface{}) error {
	if ctx.HasTimeout() {
		return nil
	}
	ctx.responseWriter.Header().Set("Content-Type", "application/json")
	ctx.responseWriter.WriteHeader(status)
	byt, err := json.Marshal(obj)
	if err != nil {
		ctx.responseWriter.WriteHeader(500)
		return err
	}
	ctx.responseWriter.Write(byt)
	return nil
}

func (ctx *Context) HTML(status int, obj interface{}, template string) error {
	return nil
}

func (ctx *Context) Text(status int, obj string) error {
	return nil
}

// #endregion

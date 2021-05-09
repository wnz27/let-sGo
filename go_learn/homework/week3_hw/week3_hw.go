/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/5/9 17:30 5月
 **/
package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func someService(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bcdefg"))
}

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	serviceMux := http.NewServeMux()
	serviceMux.HandleFunc("/a", someService)

	// 模拟中断服务
	serverOut := make(chan struct{})
	//serverOutMux := http.NewServeMux()
	serviceMux.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		serverOut <- struct{}{}
	})

	// 这里产生重复了是不是还可以优化一下？有一个类似servers 的slice?
	serviceServer := http.Server{
		Handler: serviceMux,
		Addr:    ":8080",
	}

	//shutdownServer := http.Server{
	//	Handler: serverOutMux,
	//	Addr:    ":8081",
	//}

	/*
		监听地址goroutine 退出后, 后面goroutine 都会随之退出
	 */
	g.Go(
		func() error {
			return serviceServer.ListenAndServe()
		})
	/*
		退出后, 其他 都会随之退出
	*/
	//g.Go(
	//	func() error {
	//		return shutdownServer.ListenAndServe()
	//	})

	/*
	退出时，调用了 shutdown, context 会做相应取消
	 */
	g.Go(func() error {
		select {
		case <-ctx.Done():
			log.Println("errgroup exit!")
		case <-serverOut:
			log.Println("server out!")
		}

		timeoutCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		log.Println("shutting down server!")
		return serviceServer.Shutdown(timeoutCtx)
	})

	/*
	捕获到 os 退出信号将会退出
	 */
	g.Go(func() error {
		quit := make(chan os.Signal, 0)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-ctx.Done():
			return ctx.Err()
		case sig := <-quit:
			return errors.Errorf("get os signal: %v", sig)
		}
	})

	fmt.Printf("errgroup done!!!: %+v\n", g.Wait())
}


/*
todo 多监听了一个端口号就hang 死了， Wait 产生了永久等待在这个位置： shutting down server!, 后面再找下问题。。。。。蛋疼。。。找到了问题
也解决了但是还是感觉有点儿蠢。。。
【见： demo_2.go】
 */

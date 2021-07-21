/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/7/21 4:50 下午
* Description:
 */
package main

import (
	"context"
	"github.com/go-kit/kit/transport/http"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	svr := http.NewServer()
	// http server
	g.Go(
		func() error {
		fmt.Println("http")
		go func() {
			<-ctx.Done()
			fmt.Println("http ctx done")
			svr.Shutdown(context.TODO())
		}()
		return svr.Start()
	})

	// signal
	g.Go(func() error {
		exitSignals := []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT} // SIGTERM is POSIX specific
		sig := make(chan os.Signal, len(exitSignals))
		signal.Notify(sig, exitSignals...)
		for {
			fmt.Println("signal")
			select {
			case <-ctx.Done():
				fmt.Println("signal ctx done")
				return ctx.Err()
			case <-sig:
				// do something
				return nil
			}
		}
	})

	// inject error
	g.Go(func() error {
		fmt.Println("inject")
		time.Sleep(time.Second)
		fmt.Println("inject finish")
		return errors.New("inject error")
	})

	err := g.Wait() // first error return
	fmt.Println(err)
}

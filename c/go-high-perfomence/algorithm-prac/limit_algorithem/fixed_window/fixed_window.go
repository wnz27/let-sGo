/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/7/22 2:32 下午
* Description:
 */
package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func tryReqIn() bool {
	return false
}

func main() {
	//d := time.Duration(10)
	//fmt.Println(d)
	g, ctx := errgroup.WithContext(context.Background())
	var ticker *time.Ticker = time.NewTicker(1 * time.Second)
	g.Go(
		func() error {
			for t := range ticker.C {
				fmt.Println("Tick at", t)
				// Do some thing
			}
			exitSignals := []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT} // SIGTERM is POSIX specific
			sig := make(chan os.Signal, len(exitSignals))
			signal.Notify(sig, exitSignals...)
			for {
				select {
				case <-ctx.Done():
					fmt.Println("signal ctx done")
					return ctx.Err()
				case <-sig:
					// do something
					fmt.Println("exit ticker")
					ticker.Stop()
					return nil
				}
			}
		})
	err := g.Wait() // first error return
	fmt.Printf("Ticker stopped: %s", err)
}

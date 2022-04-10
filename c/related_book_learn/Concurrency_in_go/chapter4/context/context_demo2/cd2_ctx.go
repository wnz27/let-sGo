/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/24 23:48 8月
 **/
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	// 1、这里 main 用 context.Background() 创建一个 Context 并用 context.WithCancel 包装它以允许取消。
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printGreeting(ctx); err != nil {
			fmt.Printf("cannot print greeting: %v\n", err)
			// 2、这一行上，如果从打印语问候语返回错误，main将取消这个context。
			cancel()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printFarewell(ctx); err != nil {
			fmt.Printf("cannot print farewell: %v\n", err)
		}
	}()
	wg.Wait()
}

func printGreeting(ctx context.Context) error {
	greeting, err := genGreeting(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%s world!\n", greeting)
	return nil
}

func printFarewell(ctx context.Context) error {
	farewell, err := genFarewell(ctx)
	if err != nil {return err}
	fmt.Printf("%s world!\n", farewell)
	return nil
}

func genGreeting(ctx context.Context) (string, error) {
	// 3、这里genGreeting 用context.WithTimeout 包装它的Context。这将
	// 在1s 后自动取消返回的 context，从而取消它传递该 context 的任何子函数，即语言环境。
	ctx, cancel := context.WithTimeout(ctx, 1 * time.Second)
	defer cancel()
	switch locale, err := locale(ctx); {
	case err != nil:
		return "", err
	case locale == "EN/US":
		return "hello", nil
	}
	return "", fmt.Errorf("unsupported locale")
}

func genFarewell(ctx context.Context) (string, error) {
	switch locale, err := locale(ctx); {
	case err != nil:
		return "", err
	case locale == "EN/US":
		return "goodbye", nil
	}
	return "", fmt.Errorf("unsupported locale")
}

func locale(ctx context.Context) (string, error) {
	select {
	case <-ctx.Done():
		// 这一行返回为什么Context 被取消对的原因。该错误将会一直弹出到main， 这会导致取消。
		return "", ctx.Err()
	case <-time.After(30 * time.Second):
	}
	return "EN/US", nil
}


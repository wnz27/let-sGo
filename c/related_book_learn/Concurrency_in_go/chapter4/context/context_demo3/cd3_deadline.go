/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/25 01:30 8月
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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := printGreeting(ctx); err != nil {
			fmt.Printf("cannot print greeting: %v\n", err)
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
	ctx, cancel := context.WithTimeout(ctx, 2 * time.Second)
	defer cancel()
	switch locale, err := locale(ctx); {
	case err != nil:
		return "", err
	case locale == "EV/US":
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
	// 这里我们检查我们的context 是否提供了超时时间。如果确实如此，并且我们的系统时钟已经超过截止时间，
	// 那么我们只会返回context 包中定义的特定错误，即 DeadlineExceeded
	fmt.Println("XXXXX")
	if deadline, ok := ctx.Deadline(); ok {  // Todo 没看懂, 暂时理解 这里这个 ok 会一直check 到没到 deadline
		fmt.Println("III8IIIIII")
		if deadline.Sub(time.Now().Add(1 * time.Second)) <= 0 {
			return "", context.DeadlineExceeded
		}
	}
	select {
	case <- ctx.Done():
		return "", ctx.Err()
	case <-time.After(1 * time.Minute):
	}
	return "EN/US", nil
}


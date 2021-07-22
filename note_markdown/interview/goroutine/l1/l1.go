/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/7/21 4:38 下午
* Description:
 */
package main

import (
	"fmt"
	"errors"
	"sync"
	"time"
)

func main() {
	gerrors := make(chan error)
	wgDone := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		wg.Done()
	}()

	go func() {
		err := returnError()
		if err != nil {
			gerrors <- err
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(wgDone)
	}()

	select {
	case <-wgDone:
		break
	case err := <-gerrors:
		close(gerrors)
		fmt.Println(err)
	}

	time.Sleep(time.Second)
}

func returnError() error {
	return errors.New("煎鱼报错了...")
}

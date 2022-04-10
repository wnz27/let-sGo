/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/9/15 9:27 下午
* Description:
 */
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	/*
	 Reading atomics safely while they are being updated is also possible,
	using functions like atomic.LoadUint64.
	 */

	var ops uint64

	var wg sync.WaitGroup

	for i := 0; i < 50; i ++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c ++ {
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("ops: ", ops)
}

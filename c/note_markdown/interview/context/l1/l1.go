/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/7/20 4:06 下午
* Description:
 */
package main

import (
	"context"
	"fmt"
	"time"
)

const shortDuration = 1 * time.Millisecond

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), shortDuration)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("脑子进煎鱼了")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

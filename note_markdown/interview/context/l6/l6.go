/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/7/20 4:29 下午
* Description:
 */
package main

import (
	"context"
	"fmt"
)

func main() {
	type favContextKey string
	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}

	k := favContextKey("脑子进")
	ctx := context.WithValue(context.Background(), k, "煎鱼")

	f(ctx, k)
	f(ctx, favContextKey("小咸鱼"))
}

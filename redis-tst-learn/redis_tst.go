/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/6/19 12:55 上午
* Description:
 */
package redis_tst_learn

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"unsafe"
)

func setValue() {
	var n2 int64 = 10
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6377",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	var ctx = context.Background()

	err := rdb.Set(ctx, "key1", n2, 0).Err()
	if err != nil {
		panic(err)
	}


	fmt.Printf("\n n2 的类型 %T n2占中的字节数是 %d \n", n2, unsafe.Sizeof(n2))

	//val, err := rdb.Get(ctx, "key").Result()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("key", val)

	//rdb.Set(ctx, "t1", , 300)
}

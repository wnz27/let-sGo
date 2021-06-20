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
	"strconv"
	"unsafe"
)

// groupNum 组数 dataSize 数据数量
func setValue(groupNum int, dataSize int) {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	var ctx = context.Background()

	var keys []string
	for i := 0; i < groupNum; i ++{
		currI := int64(i)
		intStr := strconv.FormatInt(currI, 10)
		keys = append(keys, "key" + intStr)
	}

	var slice1 []int = make([]int, dataSize)

	fmt.Printf("slice1, 占用字节:%d \n", unsafe.Sizeof(slice1))


	var n2 int64 = 10 // 8字节
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

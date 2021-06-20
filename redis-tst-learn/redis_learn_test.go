/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/6/19 12:56 上午
* Description:
 */
package redis_tst_learn

import (
	"context"
	"fmt"
	"testing"
	"github.com/shirou/gopsutil/mem"
	"unsafe"
	"github.com/go-redis/redis/v8"
)

func TestRedisSetGetCapacityMemInfo(t *testing.T) {
	displayMemB := func(t *testing.T) {
		t.Helper()  // ！！！！！会打出具体失败位置
		v, _ := mem.VirtualMemory()
		fmt.Printf("Before ---> Total: %v, Available: %v, UsedPercent:%f%%\n", v.Total, v.Available, v.UsedPercent)
		//fmt.Println(v)
	}

	displayMemA := func(t *testing.T) {
		t.Helper()  // ！！！！！会打出具体失败位置
		v, _ := mem.VirtualMemory()
		fmt.Printf("After ---> Total: %v, Available: %v, UsedPercent:%f%%\n", v.Total, v.Available, v.UsedPercent)
		//fmt.Println(v)
	}

	setValue2Redis := func(t *testing.T, value interface{}) {
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

		//val, err := rdb.Get(ctx, "key").Result()
		//if err != nil {
		//	panic(err)
		//}
		//fmt.Println("key", val)

		//rdb.Set(ctx, "t1", , 300)
		fmt.Printf("\n n2 的类型 %T n2占中的字节数是 %d \n", n2, unsafe.Sizeof(n2))
	}

	t.Run("before set", func(t *testing.T) {
		displayMemB(t)
	})

	t.Run("set value 2 redis", func(t *testing.T) {
		setValue2Redis(t, "abc")
	})

	t.Run("after set", func(t *testing.T) {
		displayMemA(t)
	})
}




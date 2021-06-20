/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/6/19 12:55 上午
* Description:
 */
package redis_tst_learn

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
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

	var longStr string
	for i := 0; i < dataSize; i ++ {
		currI := int64(i)
		longStr += strconv.FormatInt(currI, 10)
	}

	// 存json数据
	imap := map[string]string{}
	for i := 0; i < dataSize; i ++ {
		currI := int64(i)
		intStr := strconv.FormatInt(currI, 10)
		k := "key" + intStr
		imap[k] = longStr
	}

	// 将map转换成json数据
	v1, _ := json.Marshal(imap)

	fmt.Printf("imap, 占用字节:%d \n", unsafe.Sizeof(v1))

	for _, key := range keys {
		//fmt.Println(key)
		_ = rdb.Set(ctx, key, v1, time.Second * 300).Err()
	}
}

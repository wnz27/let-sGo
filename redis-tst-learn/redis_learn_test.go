/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/6/19 12:56 上午
* Description:
 */
package redis_tst_learn

import (
	"fmt"
	"github.com/shirou/gopsutil/mem"
	"testing"
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
		setValue(1000, 10000)

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




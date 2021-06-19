/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/6/19 12:56 上午
* Description:
 */
package redis_tst_learn

import (
	"fmt"
	"testing"
	"github.com/shirou/gopsutil/mem"
)

func TestBefore(t *testing.T) {
	displayMem := func(t *testing.T) {
		t.Helper()  // ！！！！！会打出具体失败位置
		v, _ := mem.VirtualMemory()
		fmt.Printf("Before ---> Total: %v, Available: %v, UsedPercent:%f%%\n", v.Total, v.Available, v.UsedPercent)
		//fmt.Println(v)
	}

	t.Run("before set", func(t *testing.T) {
		displayMem(t)
	})
}

func TestAfter(t *testing.T) {
	displayMem := func(t *testing.T) {
		t.Helper()  // ！！！！！会打出具体失败位置
		v, _ := mem.VirtualMemory()
		fmt.Printf("After ---> Total: %v, Available: %v, UsedPercent:%f%%\n", v.Total, v.Available, v.UsedPercent)
		//fmt.Println(v)
	}

	t.Run("after set", func(t *testing.T) {
		displayMem(t)
	})
}



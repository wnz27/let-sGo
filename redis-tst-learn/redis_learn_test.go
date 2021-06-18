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

func TestHelloTo(t *testing.T) {
	displayMem := func(t *testing.T) {
		t.Helper()  // ！！！！！会打出具体失败位置
		v, _ := mem.VirtualMemory()
		fmt.Printf("Total: %v, Available: %v, UsedPercent:%f%%\n", v.Total, v.Available, v.UsedPercent)
		fmt.Println(v)
	}

	t.Run("before set", func(t *testing.T) {
		displayMem(t)
	})
}




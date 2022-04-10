/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/14 02:04 8月
 **/
package main

import (
	"bytes"
	"fmt"
	"sync"
)

func main() {
	printData := func(wg *sync.WaitGroup, data []byte) {
		defer wg.Done()

		var buff bytes.Buffer
		for _, b := range data {
			fmt.Fprintf(&buff, "%c", b)
		}
		fmt.Println(buff.String())
	}

	var wg sync.WaitGroup
	wg.Add(2)
	data := []byte("golang")
	go printData(&wg, data[:3])  // 传入前三个字节的切片
	go printData(&wg, data[3:])  // 传入最后三个字节的切片
	wg.Wait()
}
